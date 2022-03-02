package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/atultherajput/go_crash_course/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

var id = flag.String("id", "7", "Ftech book by id")

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := protos.NewBooksClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetBook(ctx, &protos.BookIdRequest{Id: *id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("data = %s", r)
}
