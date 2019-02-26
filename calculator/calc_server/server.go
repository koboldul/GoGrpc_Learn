package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../calcpb"
	grpc "google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("calc invoked with request %v \n", req)

	firstNo := req.GetFirstNumber()
	secondNo := req.GetSecondNumber()

	result := firstNo + secondNo

	res := &calcpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:55557")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcSvcServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error to serve: %v", err)
	}
}
