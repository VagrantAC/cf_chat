package main

import (
	"context"
	pb "demo/proto"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)
var (
  port = flag.Int("port", 6789, "The server port")
)

type server struct {
  pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.EchoResponse{
		Message: "Hello" + in.GetMessage(),
	},
	nil
}

func main() {
  flag.Parse()
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterEchoServiceServer(s, &server{})
  log.Printf("server listening at %v", lis.Addr())
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
