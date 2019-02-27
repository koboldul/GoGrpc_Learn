package common

import "log"

func IsSuccess(err error, msg string) bool {
	isOk := true

	if err != nil {
		log.Fatalf("%v %v", msg, err)
		isOk = false
	}

	return isOk
}
