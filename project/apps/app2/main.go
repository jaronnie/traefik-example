package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/api/v1/echo", func(ctx *gin.Context) {
		ctx.String(200, "hello, app2")
	})

	e.Run(":9000")
}
