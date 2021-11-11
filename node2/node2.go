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
		c.RequestToken(context.Background(), &proto.Node{Id: id}) // sync operation
		log.Printf("Node %d entered CS", id)
		time.Sleep(2 * time.Second) // In CS
		log.Printf("Node %d left CS", id)
		c.ReturnToken(context.Background(), &proto.Token{From: id}) // sync operation
		time.Sleep(10 * time.Second)
	}

}
