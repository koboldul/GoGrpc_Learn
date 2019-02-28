package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	doBidiStream(c, []int32{34, 37, 59, 7, 99, 0, -99})
}

func doClientStream(c calcpb.CalcSvcClient, a []int32) {
	stream, err := c.ComputeAverage(context.Background())

	if common.IsSuccess(err, "Error while calling compute avg") {
		for _, n := range a {
			rq := &calcpb.ComputeAverageRequest{
				Number: n,
			}
			stream.Send(rq)
		}

		rsp, err := stream.CloseAndRecv()
		if common.IsSuccess(err, "Error finishing request") {
			fmt.Printf("Response geeting: %v \n", rsp.GetAverage())
		}
	}
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

func doBidiStream(c calcpb.CalcSvcClient, a []int32) {
	stream, err := c.ComputeMax(context.Background())

	if !common.IsSuccess(err, "Error while calling compute avg") {
		return
	}

	waitc := make(chan struct{})

	go func() {
		defer stream.CloseSend()

		for _, n := range a {
			rq := &calcpb.ComputeMaxRequest{
				Number: n,
			}
			err := stream.Send(rq)
			if !common.IsSuccess(err, fmt.Sprintf("Error sending number request %d", n)) {
				break
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer close(waitc)

		for {
			rsp, err := stream.Recv()

			switch {
			case err == io.EOF:
				break
			case !common.IsSuccess(err, "Error receiving response"):
				break
			default:
				fmt.Printf("Max so far is: %v\n", rsp.GetMax())
			}
		}
	}()

	<-waitc
}
