package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

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
	doUnaryWithDeadline(c)
}

func doBidiStreaming(c greetx.GreetServiceClient) {
	stream, err := c.GreetAll(context.Background())

	requests := []*greetx.GreetAllRequest{
		&greetx.GreetAllRequest{Greeting: &greetx.Greeting{FirstName: "xxx1", SecondName: "yyy1"}},
		&greetx.GreetAllRequest{Greeting: &greetx.Greeting{FirstName: "xxx2", SecondName: "yyy1"}},
		&greetx.GreetAllRequest{Greeting: &greetx.Greeting{FirstName: "xxx3", SecondName: "yyy1"}},
		&greetx.GreetAllRequest{Greeting: &greetx.Greeting{FirstName: "xxx4", SecondName: "yyy1"}},
		&greetx.GreetAllRequest{Greeting: &greetx.Greeting{FirstName: "xxx5", SecondName: "yyy1"}},
	}

	if common.IsSuccess(err, "Error whilw calling greeting all") {

		waitc := make(chan struct{})

		go func() {
			for _, rq := range requests {
				fmt.Printf("Sending message: %v\n", rq)
				err := stream.Send(rq)
				if !common.IsSuccess(err, fmt.Sprintf("Error sending request for %v ", rq)) {
					break
				}
				time.Sleep(time.Second)
			}
			stream.CloseSend()
		}()

		go func() {
			for {
				rsp, err := stream.Recv()
				if err == io.EOF {
					break
				} else {
					if common.IsSuccess(err, "Error finishing request") {
						fmt.Printf("Response geeting: %v \n", rsp.GetResult())
					} else {
						break
					}
				}
			}
			close(waitc)
		}()

		<-waitc
	}
}

func doClientStreaming(c greetx.GreetServiceClient) {
	stream, err := c.LongGreet(context.Background())

	requests := []*greetx.LongGreetRequest{
		&greetx.LongGreetRequest{Greeting: &greetx.Greeting{FirstName: "xxx1", SecondName: "yyy1"}},
		&greetx.LongGreetRequest{Greeting: &greetx.Greeting{FirstName: "xxx2", SecondName: "yyy1"}},
		&greetx.LongGreetRequest{Greeting: &greetx.Greeting{FirstName: "xxx3", SecondName: "yyy1"}},
		&greetx.LongGreetRequest{Greeting: &greetx.Greeting{FirstName: "xxx4", SecondName: "yyy1"}},
		&greetx.LongGreetRequest{Greeting: &greetx.Greeting{FirstName: "xxx5", SecondName: "yyy1"}},
	}

	if common.IsSuccess(err, "Error whilw calling LongGreet") {
		for _, rq := range requests {
			stream.Send(rq)
		}

		rsp, err := stream.CloseAndRecv()
		if common.IsSuccess(err, "Error finishing request") {
			fmt.Printf("Response geeting: %v \n", rsp.GetResult())
		}
	}
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

func doUnaryWithDeadline(c greetx.GreetServiceClient) {
	rq := &greetx.GreetWithDeadlineRequest{
		Greeting: &greetx.Greeting{
			FirstName:  "dedline",
			SecondName: "dddddd",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rsp, err := c.GreetWithDeadline(ctx, rq)

	expectedErrors := map[codes.Code]string{
		codes.DeadlineExceeded: "Timeout while greeting!",
	}
	if common.IsGrpcSuccess(err, "Error occured while calling greet", expectedErrors) {
		log.Printf("Response, %v \n", rsp.Result)
	}
}
