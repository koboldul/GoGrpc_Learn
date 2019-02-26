package main

import (
	"context"
	"fmt"
	"log"

	greetx "../greetpb"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client")

	conn, err := grpc.Dial("localhost:55555", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer conn.Close()

	c := greetx.NewGreetServiceClient(conn)

	doUnary(c)
}

func doUnary(c greetx.GreetServiceClient) {
	rq := &greetx.GreetRequest{
		Greeting: &greetx.Greeting{
			FirstName:  "tyty",
			SecondName: "yyy",
		},
	}

	rsp, err := c.Greet(context.Background(), rq)

	if err != nil {
		log.Fatalf("Error occured while calling greet: %v", err)
	}

	log.Printf("Response, %v", rsp.Result)
}
