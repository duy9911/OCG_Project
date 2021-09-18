package handler

import "fmt"

func ErorLog(msg string, err error) {
	if err != nil {
		fmt.Printf(msg, err)
	}
}
