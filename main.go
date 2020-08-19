package main

import (
	"context"
	"log"
	"time"

	pb "github.com/qiitacopy/article/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	testGetByID()

}

func testGetByID (){
	// クライアント(テスト)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewArticleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	articleID := int32(1)
	r, err := c.GetByID(ctx, &pb.ArticleID{Id: articleID})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Text)
}