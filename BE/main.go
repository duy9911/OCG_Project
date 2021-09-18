package main

import (
	"backend/model"
	"backend/router"
	"fmt"
)

func main() {
	fmt.Println("start")
	model.Create()
	router.HandleRequests()
}
