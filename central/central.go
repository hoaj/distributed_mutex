package main

import (
	"context"
	"errors"
	"log"
	"net"

	proto "github.com/hoaj/distributed_mutex/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedCentralServer
	queue chan int64
	inCS  chan int64
}

func (s *Server) requestToken(ctx context.Context, node *proto.Node) (*proto.Token, error) {
	s.queue <- node.GetId()
	s.inCS <- <-s.queue
	return &proto.Token{}, errors.New("requestToken failed")
}

func (s *Server) returnToken(ctx context.Context, node *proto.Token) (*proto.Ack, error) {
	<-s.inCS
	return &proto.Ack{}, errors.New("requestToken failed")
}

func newServer() *Server {
	return &Server{
		inCS:  make(chan int64),
		queue: make(chan int64, 3),
	}
}

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCentralServer(grpcServer, newServer())
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
