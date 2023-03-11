package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接数据库
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
		return
	}
	defer c.Close()
	// 操作数据库
	reply, err := c.Do("set","zstu","candy")

	// 利用回复助手去调用
	r, err := redis.String(reply, err)
	fmt.Println(r, err)
}
