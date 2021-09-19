package main

import (
	"backend/model"
	"backend/router"
	"fmt"

	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {
	fmt.Println("start")
	model.Create()
	router.HandleRequests()
}
