package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"../../common"

	"../greetpb"
	grpc "google.golang.org/grpc"
)

const port = 55555

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet invoked with request %v \n", req)

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("Greet invoked with request %v \n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("Client canceled the request")
			return nil, status.Error(codes.Canceled, "Timeout!")
		}
		time.Sleep(2 * time.Second)
	}

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes invoked with request %v \n", req)

	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello: %s with number: %d", firstName, i)

		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	var sb strings.Builder

	for {
		rq, error := stream.Recv()
		if error == io.EOF {
			break
		}
		if common.IsSuccess(error, "Receiving part of request failed!") {
			sb.WriteString(fmt.Sprintf("%v \n", rq.GetGreeting().GetFirstName()))
		}
	}

	return stream.SendAndClose(&greetpb.LongGreetResponse{
		Result: sb.String(),
	})
}

func (*server) GreetAll(stream greetpb.GreetService_GreetAllServer) error {
	fmt.Println("Received all request")

	for {
		rq, err := stream.Recv()

		fmt.Printf("Received request %v \n", rq)

		if err == io.EOF {
			break
		}

		if !common.IsSuccess(err, "Error getting rq") {
			return err
		}

		firstName := rq.GetGreeting().GetFirstName()
		result := fmt.Sprintf("Hello %v !", firstName)

		err = stream.Send(&greetpb.GreetAllResponse{
			Result: result,
		})

		if !common.IsSuccess(err, "Error sending the response") {
			return err
		}
	}

	return nil
}

func main() {

	fmt.Printf("Waiting requests on port %d ... \n", port)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if !common.IsSuccess(err, "Error:") {
		return
	}

	certFile := "../../env/server.crt"
	keyFile := "../../env/server.pem"

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	if common.IsSuccess(err, "Failed loading certificates: ") {
		opts := grpc.Creds(creds)
		s := grpc.NewServer(opts)

		greetpb.RegisterGreetServiceServer(s, &server{})

		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error to serve: %v", err)
		}
	}

}
