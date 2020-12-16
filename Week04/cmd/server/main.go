package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	pb "week04/api/article/v1"
	"week04/internal/pkg/grpc"
	"week04/internal/service"

	"golang.org/x/sync/errgroup"
)

const (
	addr = ":9999"
)

func main() {
	// init service api
	uc := InitArticleUseCase()
	svs := service.NewArticleService(uc)

	// register grpc service
	s := grpc.NewServer(addr)
	pb.RegisterArticleServer(s.Server, svs)

	// context
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// start grpc server
	g.Go(func() error {
		return s.Start(ctx)
	})

	// catch signals
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	// wait stop
	if err := g.Wait(); err != nil {
		log.Printf("error: %v", err)
	}
}
