package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	requests := 0

	router.GET("/pingpong", func(ctx *gin.Context) {
		ctx.String(200, "pong "+fmt.Sprint(requests))
		if ctx.Request.Header.Get("User-Agent") != "Go-http-client/1.1" {
			requests++
		}
	})

	router.Run("0.0.0.0:8080")
}
