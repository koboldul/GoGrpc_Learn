package main

import (
	"context"
	"log"

	calcpb "../calcpb"
)

func doUnary(c calcpb.CalcSvcClient, a int32, b int32) {
	rq := &calcpb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	}

	rsp, err := c.Sum(context.Background(), rq)rq := &calcpb.SumRequest{
		FirstNumber:  a,
		SecondNumber: b,
	}

	rsp, err := c.Sum(context.Background(), rq)

	if err != nil {
		log.Fatalf("Error occured while calling sum op: %v", err)
	}

	log.Printf("Result of %d + %d = %v", a, b, rsp.SumResult)
}
