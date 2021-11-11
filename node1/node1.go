package main

import (
	"context"
	"log"
	"time"

	proto "github.com/hoaj/distributed_mutex/proto"
	"google.golang.org/grpc"
)

var (
	id int64 = 1
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
		time.Sleep(2 * time.Second)
		log.Printf("Node %d left CS", id)
		c.ReturnToken(context.Background(), &proto.Token{}) // sync operation
	}

}
