/**
 * @Author: vincent
 * @Description:
 * @File:  server
 * @Version: 1.0.0
 * @Date: 2020/12/16 18:15
 */

package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server        //grpc server
	address      string //listen addr
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	log.Println("grpc server start: ", s.address)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
		log.Println("server graceful stop")
	}()

	// serve
	return s.Serve(listener)
}

func NewServer(addr string) *Server {
	server := grpc.NewServer()
	return &Server{
		Server:  server,
		address: addr,
	}
}
