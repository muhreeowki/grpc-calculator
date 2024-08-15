package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/muhreeowki/calculator-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("server", "localhost:8080", "The server address in the format of host:port")
	flag.Parse()

	// Enable TLS for the client
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// Create a connection to the server
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{1, 2, 3, 4, 5},
	})
	if err != nil {
		log.Fatalf("failed to call Sum: %v", err)
	}
	fmt.Println(res)
}
