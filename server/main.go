package main

import (
    "context"
    "fmt"
    "log"
    "net"
    //"os"

    "google.golang.org/grpc"
    pb "../pingpong"
)

const (
    port = 50051
)

// struct to implement the pingpong server
type server struct{}

// PingPongBackend implements pingpong.PingPongService
func (s *server) PingPongBackend(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
    log.Printf("Received: %v", in.Ping)
    return &pb.Pong{Pong: "Hello " + in.Ping + "!"}, nil
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterPingPongServiceServer(s, &server{})
    log.Printf("Listening on port: %d", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
