package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/config"
	"github.com/ms-url-shortner/app/logging"
	"github.com/ms-url-shortner/app/route"
)

const listenPort = ":4242"

func main() {
	gin := gin.Default()

	config.Initialize()
	log.Println("config Initialized Successfully")

	logging.Initialize(config.Config)
	log.Println("log initialised successfully")

	lw := logging.LogForFunc()

	route.Initialize(gin)
	lw.Debug("routes initialized successfully")

	if err := gin.Run(listenPort); err != nil {
		lw.Fatalf("gin engine failed to run", err.Error())
	}
}
