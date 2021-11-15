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

func (s *Server) Dequeue() {
	for {
		mu.Lock()                 // Only dequeues if unlocked.
		NodeForQueue := <-s.queue // remove from queue
		id := <-NodeForQueue      // Let RequestToken function move on
		log.Printf("Node: %v got dequeued\n", id)
	}
}

func (s *Server) RequestToken(ctx context.Context, node *proto.Node) (*proto.Token, error) {
	NodeForQueue := make(chan int) // Representation of a node for the queue
	log.Printf("Node: %v got enqueued\n", node.GetId())
	s.queue <- NodeForQueue           // Move "Node" into queue channel
	NodeForQueue <- int(node.GetId()) // Sync call. Waiting for release in function "Dequeue".
	time.Sleep(time.Second)           // used to better see that only one node is in CS when using logs
	log.Printf("Node: %d just entered the CS", node.GetId())
	return &proto.Token{From: node.GetId()}, nil
}

func (s *Server) ReturnToken(ctx context.Context, node *proto.Token) (*proto.Ack, error) {
	mu.Unlock()
	log.Printf("Node: %d just left the CS", node.GetFrom())
	return &proto.Ack{From: node.GetFrom()}, nil
}

func newServer() *Server {
	s := &Server{
		queue: make(chan (chan int), 3), // size of queue
	}
	go s.Dequeue()
	return s
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
