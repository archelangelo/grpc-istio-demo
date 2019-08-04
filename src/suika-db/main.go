package main

import (
	"os"
	"log"

	"google.golang.org/grpc"
	pb "github.com/archelangelo/grpc-istio-demo/src/proto"
)

const (
	port = os.Getenv("SUIKA_DB_PORT")
)

type server struct{}

func (s *server) Lookup(id *pb.Id) error {
	log.Printf("Received: %v", id.Id)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v", err)
	}
	return &pb.Document{"ZZWZ", 26, "Baoshan, Shanghai"}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSuikaServer(s, &server{})
	log.Printf("Listening on port: %d", port)
	if err := s.Server(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}