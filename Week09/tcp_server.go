package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	PORT              = ":8888"
	MSG_CMD_CHAT      = "chat"
	MSG_CMD_BYE       = "bye"
	MSG_CMD_SEND_GIFT = "send gift"
)

func main() {
	// listen
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("TCP SERVER LISTEN ON %s", PORT)

	// with cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handle signal: SIGINT, SIGTERM
	go func() {
		// prepare
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		// catch signal
		sig := <-signals
		log.Printf("signal caught: %s, ready to quit...", sig.String())

		// notify cancel and close
		cancel()
		listener.Close()
	}()

	// waitGroup establish
	wg := &sync.WaitGroup{}
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		// handle conn
		wg.Add(1) //waitGroup add
		go handleConn(conn, wg, ctx)
	}

	// waitGroup wait
	wg.Wait()

	log.Print("All CONN CLOSE, TCP SERVER QUIT")
}

func handleConn(conn net.Conn, w *sync.WaitGroup, ctx context.Context) {
	// main waitGroup
	defer w.Done()

	remoteAddr := conn.RemoteAddr()

	log.Printf("%s conn accepted", remoteAddr)
	defer log.Printf("%s conn closed", remoteAddr)

	ch := make(chan []byte, 1024)

	// sub waitGroup
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// handle read and write
	go handleRead(conn, ch, wg, ctx)
	go handleWrite(conn, ch, wg, ctx)

	// wait handleRead and handleWrite closed
	wg.Wait()
}

func handleRead(conn net.Conn, ch chan<- []byte, w *sync.WaitGroup, ctx context.Context) {
	remoteAddr := conn.RemoteAddr()
	defer log.Printf("%s read exit", remoteAddr)
	defer w.Done()

	// 通知 write 协程退出
	defer close(ch)

	rd := bufio.NewReader(conn)

	for {
		// get msg
		msg, _, err := rd.ReadLine()
		if err != nil {
			return
		}

		// dispatch cmd
		switch string(msg) {
		case MSG_CMD_CHAT:
			//todo chat
			log.Printf("%s receive MSG_CMD_CHAT", remoteAddr)
			return
		case MSG_CMD_BYE:
			log.Printf("%s receive MSG_CMD_BYE", remoteAddr)
			return
		case MSG_CMD_SEND_GIFT:
			log.Printf("%s receive MSG_CMD_SEND_GIFT", remoteAddr)
			return
		}

		// not cmd, dispatch
		select {
		case ch <- msg:
		case <-ctx.Done():
			return
		}
	}
}

func handleWrite(conn net.Conn, ch <-chan []byte, wg *sync.WaitGroup, ctx context.Context) {
	connAddr := conn.RemoteAddr()

	defer log.Printf("%s write exit", connAddr)

	defer wg.Done()

	// write 协程退出关闭 client 连接
	defer conn.Close()

	wr := bufio.NewWriter(conn)

	for {
		select {
		case line, ok := <-ch:
			if !ok {
				return
			}

			if len(line) <= 0 {
				continue
			}

			// echo msg
			msg := fmt.Sprintf("echo %s\n", string(line))
			wr.WriteString(msg)
			wr.Flush()
		case <-ctx.Done():
			return
		}
	}
}
