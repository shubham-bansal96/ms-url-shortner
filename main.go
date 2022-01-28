package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/config"
	"github.com/ms-url-shortner/app/route"
)

const listenPort = ":4242"

func main() {
	gin := gin.Default()

	config.Initialize()
	log.Println("config Initialized Successfully")

	route.Initialize(gin)
	log.Println("route Initialized Successfully")

	if err := gin.Run(listenPort); err != nil {
		log.Fatalf("gin engine failed to run: %v", err.Error())
	}
}
