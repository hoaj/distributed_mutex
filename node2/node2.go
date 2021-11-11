package main

import (
	"context"
	"log"
	"time"

	proto "github.com/hoaj/distributed_mutex/proto"
	"google.golang.org/grpc"
)

var (
	id int64 = 2
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()
	c := proto.NewCentralClient(conn)

	for {
		r1, _ := c.RequestToken(context.Background(), &proto.Node{Id: id})
		log.Printf("Node: %d entered CS", r1.GetFrom())
		time.Sleep(4 * time.Second)
		r2, _ := c.ReturnToken(context.Background(), &proto.Token{From: id})
		log.Printf("Node: %d left CS", r2.GetFrom())
		time.Sleep(2 * time.Second) // Added 2 sec to better see switch between nodes in CS
	}

}
