package main

import (
	"context"
	pb "github.com/archelangelo/grpc-istio-demo/src/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"

)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func shakeHand(conn *grpc.ClientConn) {
	client := pb.NewPingPongClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.ShakeHand(ctx, &pb.Ping{Ping: name})
	if err != nil {
		log.Fatalf("Could not ping: %v", err)
	}
	log.Printf("Response: %s", res.Pong)
}

func streamCall(conn *grpc.ClientConn) {
	client := pb.NewPingPongClient(conn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatalf("error occurred getting client: %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	go func() {
		for i := 1; i <= 10; i++ {
			req := pb.Ping{Ping: name}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("error occurred sending: %v", err)
			}
			log.Printf("%d times sent", i)
			time.Sleep(500 * time.Millisecond)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("error occurred closing upward stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("error occurred receiving: %v", err)
			}
			log.Printf("Response: %s", res.Pong)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		close(done)
	}()

	<-done
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	shakeHand(conn)
	time.Sleep(1 * time.Second)
	streamCall(conn)
}
