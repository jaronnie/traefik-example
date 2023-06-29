package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/gateway/:service", func(ctx *gin.Context) {
		echo := ctx.Param("service")
		ctx.String(200, "hello "+echo)
	})

	e.Run(":9000")
}
