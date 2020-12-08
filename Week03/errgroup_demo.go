package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// 增加signal处理
const MAX_TIME_WAIT_SHUTDOWN = 2 * time.Second

func main() {
	// 打印pid，为发kill用
	pid := os.Getpid()
	log.Printf("process pid: %d \n", pid)

	// (1) ------- init
	ctx, cancel := context.WithCancel(context.Background())
	g, errCtx := errgroup.WithContext(ctx)
	svr := http.Server{Addr: "127.0.0.1:8080"} // handler 暂时省略

	// (2) ------- 3个go的func

	// 2.1 第一个go：启动http
	g.Go(func() error {
		log.Println("[HTTP SERVER] Server ListenAndServe ready to start...")
		err := svr.ListenAndServe()
		log.Printf("[HTTP SERVER] Server ListenAndServe quit with error: %v\n", err)
		return err
	})

	// 2.2 第2个go： shutdown
	g.Go(func() error {
		<-ctx.Done()
		// 由于上文中内部调用了cancel(),所以会有Done()接受到了消息
		// returns an error or ctx.Done is closed
		// 在当前工作完成或者上下文被取消之后关闭
		log.Println("[SHUTDOWN] Server shutdown ready to start..., by ", ctx.Err())
		err := svr.Shutdown(errCtx)
		log.Printf("[SHUTDOWN] Server shutdown with error: %v\n", err)

		// 准备个最大的超时时间
		go func() {
			<-time.After(MAX_TIME_WAIT_SHUTDOWN)
			log.Fatalf("MAX_TIME_WAIT_SHUTDOWN time out, force shut down")
		}()

		return err
	})

	// 2.3 第3个go： signal
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		//取得中断信号
		select {
		case sig := <-sigs:
			log.Println("[SIGNAL] Receive signal: ", sig)
			// 通知cancel
			cancel()
			log.Println("[SIGNAL] Signal notify cancel")
		case <-ctx.Done():
			log.Println("[SIGNAL] Receive cancel", ctx.Err())
		}

		return nil
	})

	// todo: 测试主动cancel
	go func() {
		time.Sleep(time.Second * 5)
		log.Println("[SOMEBODY] invoke context cancel")
		cancel()
	}()

	// (3) ------- wait
	if err := g.Wait(); err == nil {
		log.Println("[MAIN_THREAD] Successfully server close.")
	}

	// todo：制造shutdown超时
	time.Sleep(time.Second * 5)
}

