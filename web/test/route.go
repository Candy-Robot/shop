package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context){
		ctx.String(200, "hello world")
	})

	route.Run(":9091")
}
