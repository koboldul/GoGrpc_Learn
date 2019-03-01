package common

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsSuccess(err error, msg string) bool {
	isOk := true

	if err != nil {
		log.Fatalf("%v %v", msg, err)
		isOk = false
	}

	return isOk
}

func IsGrpcSuccess(err error, msg string, cases map[codes.Code]string) bool {
	isOk := err == nil

	if !isOk {
		statusErr, ok := status.FromError(err)
		if ok {
			if m, ok := cases[statusErr.Code()]; ok {
				log.Fatalf("%v --> %v \n", statusErr, m)
			} else {
				log.Fatalf("Unexpected grps exception%v! \n", statusErr)
			}
		} else {
			log.Fatalf("%v %v", msg, err)
		}

		isOk = false
	}

	return isOk
}
