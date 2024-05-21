package main

import (
	pb "MeatCounter/MeatCounter"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 8080, "The server port")
	// port = ":8080"
)

type server struct {
	pb.UnimplementedMeatCounterServer
}

func (s *server) MeatCount(ctx context.Context, in *pb.Empty) (*pb.MeatRes, error) {
	meatCount := pb.MeatCount()
	return &pb.MeatRes{MeatCount: meatCount}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMeatCounterServer(s, &server{})

	reflection.Register(s)
	if e := s.Serve(lis); e != nil {
		panic(e)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
