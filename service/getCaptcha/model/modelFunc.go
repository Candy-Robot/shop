package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 存储图片id到redis数据库
func SaveImgCode(uuid, code string) error{
	// 连接数据库
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
		return err
	}
	defer c.Close()

	// 写数据库	---有效时间5分钟
	_, err = c.Do("setex",uuid, 60*3, code)
	// 不需要回复助手
	return err
}