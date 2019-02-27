package main

import (
	"context"
	"fmt"
	"io"
	"log"

	common "../../common"
	calcpb "../calcpb"
	grpc "google.golang.org/grpc"
)

const port = 55557

func main() {
	fmt.Println("Computing factors")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer conn.Close()

	c := calcpb.NewCalcSvcClient(conn)

	doStream(c, 34)
}

func doStream(c calcpb.CalcSvcClient, a int64) {
	rq := &calcpb.PrimeNoDecompRequest{
		Number: a,
	}

	rsp, err := c.PrimeNoDecomp(context.Background(), rq)

	if common.IsSuccess(err, "Error reading primes ") {
		for {
			if err == io.EOF {
				break
			} else {
				factor, err := rsp.Recv()
				if common.IsSuccess(err, "Error reading a factore : ") {
					log.Printf("Is divisible by %v pow %v \n", factor.GetPrimeFactor(), factor.GetPower())
				}
			}
		}
	}
}
