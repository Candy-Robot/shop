package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/gomodule/redigo/redis"
	"go-micro.dev/v4"
	"image/png"
	"net/http"
	"web/model"
	getCaptcha2 "web/proto/getCaptcha"
	"web/proto/user"
	"web/utils"
)

// 获取session信息
func GetSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	// 获取session数据
	s := sessions.Default(ctx)	// 初始化  session  对象
	userName:= s.Get("userName")
	if userName == nil{
		// 用户没有登陆
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	}else{
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		// 返回取得到的session
		var nameData struct{
			Name string `json:"name"`
		}
		nameData.Name = userName.(string)	// 做一个类型断言
		resp["data"] = nameData
	}
	ctx.JSON(http.StatusOK, resp)
}

// 退出登陆
func DeleteSession (ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Delete("userName")
	// 必须使用save保存
	err := s.Save()
	resp := make(map[string]interface{})

	if err != nil {
		resp["errno"] = utils.RECODE_IOERR	//暂用ioerr
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	}else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 获取图片验证码
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码的uuid
	uuid := ctx.Param("uuid")

	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
		)

	// 初始化客户端
	microClient := getCaptcha2.NewGetCaptchaService("getcaptcha", consulService.Client())
	// 调用远程方法
	resp, err := microClient.Call(context.TODO(), &getCaptcha2.CallRequest{Uuid: uuid})
	if err != nil {
		fmt.Println("microClient err", err)
		return
	}
	// 得到的是一个bytes 需要得到的数据反序列化
	var img captcha.Image

	json.Unmarshal(resp.Img, &img)

	// 将图片写出到浏览器
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
}

// 验证验证码 没有做获取短信验证码
func GetSmscd(ctx *gin.Context) {
	// 验证验证码
	//phone := ctx.Param("phone")

	// 拆分GET请求中的URL ——格式：资源路径?k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	// 创建容器，存储返回的数据信息
	resp := make(map[string]string)

	// 校验图片验证码
	result := model.CheckImgCode(uuid, imgCode)
	if result{
		// 校验成功，发送成功状态
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}else{
		// 校验失败，发送失败状态
		resp["errno"] = utils.RECODE_SMSERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SMSERR)
	}

	// 发送成功或者失败的结果
	ctx.JSON(http.StatusOK, resp)
}

// 发送注册信息
func PostRet(ctx *gin.Context) {
	var regData struct{
		Mobile string `json:"mobile"`
		Password string		`json:"password"`
		SmsCode string	`json:"sms_code"`
	}
	ctx.Bind(&regData)

	// 将手机号，密码传入数据库 微服务的方法
	consulService := utils.InitMicro()

	// 初始化客户端
	microClient := user.NewUserService("user", consulService.Client())
	// 调用远程方法
	resp, err := microClient.Register(context.TODO(), &user.RegRequest{
		Mobile: regData.Mobile,
		Password: regData.Password,
	})
	if err != nil {
		fmt.Println("注册用户失败，找不到远程服务")
		return
	}

	// 写给浏览器
	ctx.JSON(http.StatusOK, resp)


	/*
	// form表单 才可以采用这个方法
	//mobile := ctx.PostForm("mobile")
	//pwd := ctx.PostForm("password")
	//sms_code := ctx.PostForm("sms_code")
	//
	//fmt.Println("m=", mobile, "pwd=", pwd, "sms_code=",sms_code)
	 */
}

// 获取地域信息
func GetArea(ctx *gin.Context) {
	var areas []model.Area

	// 从缓存redis 中，获取数据
	conn := model.RedisPool.Get()
	// 使用字节切片接受
	areadata, _  := redis.Bytes(conn.Do("get", "areaData"))
	// 说明没有从redis中获取数据
	if len(areadata) == 0 {
		// 先从MySQL中获取数据
		model.GlobalConn.Find(&areas)

		// 再将数据写入redis中,存储结构体序列化后的json串
		areaBuf, _ := json.Marshal(areas)
		//conn := model.RedisPool.Get()
		//conn.Do("set", "areaData", areas)
		conn.Do("set", "areaData", areaBuf)
	} else {
		// 说明redis有数据
		json.Unmarshal(areadata, &areas)
	}

	resp := make(map[string]interface{})
	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 处理登陆业务
func PostLogin(ctx *gin.Context) {
	// 获取前端数据
	var logData struct{
		Mobile string `json:"mobile"`
		Password string `json:"password"`
	}
	ctx.Bind(&logData)
	fmt.Println(logData)
	resp := make(map[string]interface{})
	// 获取数据库的数据，查询是否和输入的数据匹配
	userName, err := model.Login(logData.Mobile,logData.Password)
	if err == nil{ // 登陆成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		// 将登陆状态，保存到session中
		s := sessions.Default(ctx)
		// 将用户名设置到session中
		s.Set("userName", userName)
		s.Save()
	}else{
		// 登陆失败
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	// 得到当前用户的信息
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	// 有可能没拿到 保证页面访问的健壮性
	if userName == nil{
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	// 在数据库中查找
	user, err := model.GetUserInfo(userName.(string))
	if err != nil{
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	var data struct{
		User_id int `json:"user_id"`
		Name string `json:"name"`
		Mobile string `json:"mobile"`
		Real_name string `json:"real_name"`
		Id_card string `json:"id_card"`
		Avatar_url string `json:"avatar_url"`
	}
	data.User_id = user.ID
	data.Name = user.Name
	data.Id_card = user.Id_card
	data.Mobile = user.Mobile
	data.Real_name = user.Real_name
	data.Avatar_url = user.Avatar_url
	resp["data"] = data
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 更新用户名
func PutUserInfo(ctx *gin.Context) {
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	// 获取当前用户名
	s := sessions.Default(ctx)
	userName := s.Get("userName").(string) // 进行断言

	// 获取新用户名
	var nameData struct{
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)

	// 更新用户名
	// 数据库更新
	err := model.UpdateUserInfo(userName, nameData.Name)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	// 更新Session数据
	s.Set("userName", nameData.Name)
	err = s.Save()
	if err != nil{
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

// 上传头像
func PostAvatar(ctx *gin.Context) {
	// 获取文件对象
	file, _ :=  ctx.FormFile("avatar")

	// 上传文件到项目中
	err := ctx.SaveUploadedFile(file, "test/"+file.Filename)
	fmt.Println(err)
}

// 实名认证
func PostUserAuth(ctx *gin.Context) {
	
}