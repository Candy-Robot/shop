package model

import (
	"crypto/md5"
	"encoding/hex"
)

// 注册用户信息，写到MySQL数据库
func RegistUser(mobile, password string) error{
	user := new(User)
	// 将name和手机号都先初始为手机号
	user.Mobile = mobile
	user.Name = mobile
	// 对密码进行md5加密
	data := []byte(password)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	user.Password_hash = hex.EncodeToString(md5Ctx.Sum(nil))

	// 将用户存入数据库
	return GlobalConn.Create(&user).Error
}
