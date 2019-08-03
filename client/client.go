package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	pb "../pingpong"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPingPongServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.PingPongBackend(ctx, &pb.Ping{Ping: name})
	if err != nil {
		log.Fatalf("Could not ping: %v", err)
	}
	log.Printf("Response: %s", res.Pong);
}
