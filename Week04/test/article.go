package main

import (
	"context"
	"log"
	"time"

	pb "week04/api/article/v1"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9999"
	title   = "vincent is coming"
	content = "Our hero, vincent, is ready to come here. We will do something for him."
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// grpc client
	c := pb.NewArticleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// call rpc
	r, err := c.PostArticle(ctx, &pb.PostArticleRequest{Title: title, Content: content})
	if err != nil {
		log.Fatalf("post article failed: %v", err)
	}
	log.Printf("post article id: %d", r.GetId())
}
