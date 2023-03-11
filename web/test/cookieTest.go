package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store, _ := redis.NewStore(10, "tcp","127.0.0.1:6379","",[]byte("secret"))
	//// name是cookie值
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(ctx *gin.Context){
		//// 设置cookie
		ctx.SetCookie("mytest","tyc",60*30,"","",true,true)
		//// 获取cookie-
		cookieVal, _ := ctx.Cookie("mytest")
		fmt.Println(cookieVal)

		//// Session，设置session数据
		//session := sessions.Default(ctx)
		/////*
		////// 设置session
		//session.Set("zstu","tyc")
		////// 修改session时，需要Save函数配合
		//session.Save()
		////*/
		//v := session.Get("zstu")
		//fmt.Println("获取session",v.(string))
		//var count int
		//v := session.Get("count")
		ctx.Writer.WriteString("test cookie")
	})
	router.Run(":9999")
}