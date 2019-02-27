package main

import (
	"context"
	"fmt"
	"net"

	"../../common"
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

func (*server) PrimeNoDecomp(req *calcpb.PrimeNoDecompRequest, stream calcpb.CalcSvc_PrimeNoDecompServer) error {
	fmt.Printf("Decompistion in primes invoked with request %v \n", req)

	//number := req.GetNumber()

	primes := []int64{2, 3, 5, 7}

	for i, p := range primes {
		res := &calcpb.PrimeNoDecompResponse{
			PrimeFactor: p,
			Power:       int64(i),
		}

		stream.Send(res)
	}

	return nil
}

const port = 55557

func main() {
	fmt.Printf("Starting local on port %v \n", port)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))

	if common.IsSuccess(err, "Error on starting listener") {
		s := grpc.NewServer()
		calcpb.RegisterCalcSvcServer(s, &server{})

		common.IsSuccess(s.Serve(lis), "Error registering the server")
	}
}
