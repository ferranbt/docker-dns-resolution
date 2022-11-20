package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/coredns/coredns/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:5566"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Grpc server running in 0.0.0.0:5566")

	grpcServer := grpc.NewServer()
	pb.RegisterDnsServiceServer(grpcServer, &service{})

	grpcServer.Serve(lis)
}

type service struct {
	pb.UnimplementedDnsServiceServer
}

func (s *service) Query(ctx context.Context, req *pb.DnsPacket) (*pb.DnsPacket, error) {
	fmt.Println("-req-")
	fmt.Println(string(req.Msg))

	return nil, fmt.Errorf("bad")
}
