package main

import (
	"fmt"
	"github.com/dmr/microservice/productservice/service"
	"github.com/dmr/microservice/productservice/dbclient"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.InitializeBucket()
}