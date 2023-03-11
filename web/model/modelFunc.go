package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var RedisPool redis.Pool

// 创建函数 初始化Redis连接池
func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle: 5,
		MaxActive: 10,
		MaxConnLifetime: 60 * 5,
		IdleTimeout: 60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

// 校验图片验证码
func CheckImgCode(uuid, imgCode string) bool {
	// 连接redis数据库
	/*
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
		return false
	}
	defer c.Close()
	*/
	conn := RedisPool.Get()
	defer conn.Close()

	// 查询redis数据
	code, err := redis.String(conn.Do("get", uuid))
	if err != nil{
		fmt.Println("c.Do err", err)
		return false
	}
	return code == imgCode
}

// 处理登陆业务，根据手机号/密码  获取用户名
func Login(mobile, pwd string) (string, error){
	var user User
	// 对参数pwd做md5哈希
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))
	err := GlobalConn.
		Where("mobile=?", mobile).Where("password_hash=?", pwd_hash).Find(&user).Error
	return user.Name, err
}

// 查询用户的信息
func GetUserInfo(userName string) (User, error){
	var user User
	err := GlobalConn.Where("name=?", userName).Find(&user).Error
	return user, err
}

// 更新用户名
func UpdateUserInfo(oldName, newName string) error {
	return GlobalConn.Model(new(User)).Where("name=?", oldName).Update("name", newName).Error
}