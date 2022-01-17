package main

import (
	"github.com/gin-gonic/gin"
	"server_automatization/handlers"
)

var r *gin.Engine

func main() {
	r = gin.Default()
	initialize()
	r.Run()
}

func initialize() {
	api := r.Group("/api")
	{
		api.GET("/details", handlers.Details)
	}
}
