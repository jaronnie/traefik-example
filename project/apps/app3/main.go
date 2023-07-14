package main

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Message string `json:"message"`
}

func main() {
	e := gin.Default()

	e.GET("/api/v1/echo/:message", func(ctx *gin.Context) {
		message := ctx.Param("message")
		ctx.String(200, "hello app3: "+message)
	})

	e.POST("/api/v1/post", func(ctx *gin.Context) {
		var d Data
		err := ctx.BindJSON(&d)
		if err != nil {
			ctx.String(500, "bind Data err: "+err.Error())
		}
		ctx.String(200, "post message: "+d.Message)
	})

	e.Run(":9000")
}
