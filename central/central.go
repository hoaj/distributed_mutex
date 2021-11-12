package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	proto "github.com/hoaj/distributed_mutex/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedCentralServer
	queue chan chan int
}

var mu sync.Mutex

func (s *Server) RequestToken(ctx context.Context, node *proto.Node) (*proto.Token, error) {
	NodeinQueue := make(chan int)
	go func() {
		s.queue <- NodeinQueue // Move "Node" into queue channel
		log.Printf("Node: %v got enqueued\n", node.GetId())
		go func() {
			for {
				mu.Lock()
				NodeinQueue := <-s.queue // remove from queue channel
				c := <-NodeinQueue       // Confirm removement from queue to release line 33.
				log.Printf("Node: %v got dequeued\n", c)
			}
		}()
	}()
	NodeinQueue <- int(node.GetId()) // Waiting in queue for release confirmation
	time.Sleep(time.Second)          // put in order to better see that the dequeue happens before entering CS.
	log.Printf("Node: %d just entered the CS", node.GetId())
	return &proto.Token{From: node.GetId()}, nil
}

func (s *Server) ReturnToken(ctx context.Context, node *proto.Token) (*proto.Ack, error) {
	mu.Unlock()
	log.Printf("Node: %d just left the CS", node.GetFrom())
	return &proto.Ack{From: node.GetFrom()}, nil
}

func newServer() *Server {
	return &Server{
		queue: make(chan (chan int), 3), // size of queue
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
