package main

import (
	"context"
	"fmt"
	"log"

	calcpb "../calcpb"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println("Computing sum")

	conn, err := grpc.Dial("localhost:55557", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer conn.Close()

	c := calcpb.NewCalcSvcClient(conn)

	doUnary(c, 100, 34)
}

func doUnary(c calcpb.CalcSvcClient, a int32, b int32) {
	rq := &calcpb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	}

	rsp, err := c.Sum(context.Background(), rq)

	if err != nil {
		log.Fatalf("Error occured while calling sum op: %v", err)
	}

	log.Printf("Result of %d + %d = %v", a, b, rsp.SumResult)
}
