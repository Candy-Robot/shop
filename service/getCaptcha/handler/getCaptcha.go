package handler

import (
	"context"
	"encoding/json"
	"getCaptcha/model"
	"github.com/afocus/captcha"

	pb "getCaptcha/proto"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {

	// 生成图片验证码
	cap := captcha.New()
	// 设置字体
	cap.SetFont("./config/comic.ttf")

	// 创建验证码 4个字符
	// 暂时不需要使用验证的str
	img, str := cap.Create(4, captcha.NUM)

	// 存储图片验证码到redis中
	err := model.SaveImgCode(req.Uuid, str)
	if err != nil{
		return err
	}
	// 将 生成的图片 进行序列化
	imgBuf, _ := json.Marshal(img)

	// 将imagebuf使用参数传出
	rsp.Img = imgBuf

	return nil
}

