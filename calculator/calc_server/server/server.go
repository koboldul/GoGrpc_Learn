package calcserver

import (
	"context"
	"fmt"
	"net"

	"../../../common"
	"../../calcpb"
	grpc "google.golang.org/grpc"
)

const port = 55557

type server struct {
}

func RunServer() {
	fmt.Printf("Starting local on port %v \n", port)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))

	if common.IsSuccess(err, "Error on starting listener") {
		s := grpc.NewServer()
		calcpb.RegisterCalcSvcServer(s, &server{})

		common.IsSuccess(s.Serve(lis), "Error registering the server")
	}
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
	fmt.Printf("Decompostion in primes invoked with request %v \n", req)
	number := req.GetNumber()
	processPrimes(number, stream.Send)

	return nil
}

func processPrimes(number int64, send func(*calcpb.PrimeNoDecompResponse) error) {
	div := int64(2)
	currentFact := []int64{}

	for number > 1 {
		if number%div == 0 {
			number = number / div
			currentFact = append(currentFact, div)
		} else {
			switch {
			case div == 2:
				div++
			default:
				div += 2
			}
		}

		l := len(currentFact)
		if l > 0 && currentFact[0] != currentFact[l-1] {
			send(composeResult(currentFact, l))
			currentFact = currentFact[l-1:]
		}
	}

	l := len(currentFact)
	if l > 0 {
		send(composeResult(currentFact, l+1))
	}
}

func composeResult(currentFact []int64, l int) *calcpb.PrimeNoDecompResponse {
	return &calcpb.PrimeNoDecompResponse{
		PrimeFactor: currentFact[0],
		Power:       int64(l - 1),
	}
}
