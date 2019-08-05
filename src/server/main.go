package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "github.com/archelangelo/grpc-istio-demo/src/proto"
)

const (
    port = 50051
)

// struct to implement the pingpong server
type server struct{}

// Backend implements pingpong.PingPong
func (s *server) ShakeHand(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
    log.Printf("Received: %v", in.Ping)
    hostname, err := os.Hostname()
    if err != nil {
        log.Fatalf("failed to get hostname: %v", err)
    }
    return &pb.Pong{Pong: "Hello " + in.Ping + "! This is " + fmt.Sprint(hostname) + "."}, nil
}

// PingPongStream implements pingpong.PingPong
func (s *server) Stream(srv pb.PingPong_StreamServer) error {
    log.Println("Stream starts")
    context := srv.Context()

    for {
        select {
        case <- context.Done():
            return context.Err()
        default:
        }

        req, err := srv.Recv()
        if err == io.EOF {
            log.Println("exit")
            return nil
        }
        if err != nil {
            log.Fatalf("error occurred receiving: %v", err);
        }

        hostname, err := os.Hostname()
        if err != nil {
            log.Fatalf("failed to get hostname: %v", err)
        }
        res := pb.Pong{Pong: "Hello " + req.Ping + "! This is " + fmt.Sprint(hostname) + "."}
        if err := srv.Send(&res); err != nil {
            log.Fatalf("error occured sending: %v", err)
        }
        log.Printf("Sent stream")
    }
}

func (s *server) CheckOut(ctx context.Context, in *pb.Id) (*pb.Document, error){
    log.Printf("Received: %v", in.Id)
    address := "localhost:" + os.Getenv("SUIKA_DB_PORT")

    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    client := pb.NewSuikaClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    res, err := client.Lookup(ctx, in)
    if err != nil {
        log.Fatalf("Could not ping: %v", err)
    }
    return res, err
}

func main() {

    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterPingPongServer(s, &server{})
    log.Printf("Listening on port: %d", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
