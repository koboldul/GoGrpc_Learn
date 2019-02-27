package main

import (
	"context"
	"fmt"
	"io"
	"log"

	common "../../common"
	greetx "../greetpb"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:55555", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer conn.Close()

	c := greetx.NewGreetServiceClient(conn)

	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c greetx.GreetServiceClient) {
	fmt.Println("Starting streaming ...")

	rq := &greetx.GreetManyTimesRequest{
		Greeting: &greetx.Greeting{
			FirstName:  "tyty",
			SecondName: "yyy",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), rq)
	if common.IsSuccess(err, "Error on greets ") {
		for {
			msg, err := resStream.Recv()
			if err == io.EOF {
				break
				//end of stream
			}
			if common.IsSuccess(err, "Error on greet ") {
				fmt.Printf("Result for many is %v \n", msg.GetResult())
			}
		}
	}
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

	log.Printf("Response, %v \n", rsp.Result)
}
