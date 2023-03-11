package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/model"
)

func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化Session对象
		s := sessions.Default(ctx)
		userName := s.Get("userName")

		if userName == nil{
			// 找不到用户名
			ctx.Abort()		// 从这里返回，不必继续执行
		}else {
			ctx.Next()		// 继续向下执行
		}
	}
}

// 添加gin框架3步骤
func main(){

	// 初始化Redis连接池
	model.InitRedis()

	// 初始化MySQL连接池
	model.InitDb()

	// 初始化路由
	router := gin.Default()

	// 初始化容器
	store, _ := redis.NewStore(10,"tcp", "127.0.0.1:6379","", []byte("secret"))

	// 使用容器	使用中间件
	router.Use(sessions.Sessions("mysession", store))

	// 路由匹配
	router.Static("/home", "view")

	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions", controller.PostLogin)

		r1.Use(LoginFilter())	// 以后的路由都不需要校验Session了，直接获取数据即可
		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)
		r1.POST("/user/avatar", controller.PostAvatar)
		r1.POST("/user/auth", controller.PostUserAuth)
	}

	// 启动运行
	router.Run(":8080")
}
