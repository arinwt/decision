package main

import (
	"fmt"
	"github.com/arinwt/decision/database"
	"github.com/arinwt/decision/controller"
)

func main() {
	fmt.Println("initializing...")
	database.Init()
	fmt.Println("starting...")
	controller.Start()
}