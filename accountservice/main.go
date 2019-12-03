package main

import (
	"fmt"
	service "github.com/ryanyogan/goblog/accountservice/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
