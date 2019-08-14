package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/archelangelo/grpc-istio-demo/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int

type server struct{}

func (s *server) Lookup(ctx context.Context, id *pb.Id) (*pb.Document, error) {
	log.Printf("Received: %v", id.Id)
	_, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v", err)
	}
	return &pb.Document{Name: "ZZWZ", Age: 26, Address: "Baoshan, Shanghai"}, nil
}

func main() {
	port, err := strconv.Atoi(os.Getenv("SUIKA_DB_PORT"))
	if err != nil {
		log.Fatalf("Failed to get port: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSuikaServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Listening on port: %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
