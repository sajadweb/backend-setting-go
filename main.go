package main

import (
	"bakend-settings/category"
	"bakend-settings/config"
	"bakend-settings/microservices/tcp"
	"fmt"
	"log"
)

func main() {
	config.LoadEnv()
	app := tcp.Serve(config.GetEnv("TCP_SERVER"))
	category.CateogyMain(app)
	fmt.Println("Start tcp:",config.GetEnv("TCP_SERVER"))
	log.Fatal(app.Start())
}
