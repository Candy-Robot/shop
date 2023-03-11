package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("333333")
		c.Abort()
		fmt.Println("5555555")
	}
}

func Test1(ctx *gin.Context)  {
	fmt.Println("1111111")
	ctx.Abort()
	fmt.Println("4444444")
}

func main() {
	router := gin.Default()
	// 使用中间件
	router.Use(Test1)
	router.Use(Logger())

	router.GET("/test", func(context *gin.Context) {
		fmt.Println("222222222")
		context.Writer.WriteString("hello world!")
	})

	router.Run(":9999")
}