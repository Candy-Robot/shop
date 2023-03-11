package handler

import (
	"context"
	"user/model"
	"user/utils"

	pb "user/proto/user"
)

type User struct{}

func (e *User) Register(ctx context.Context, req *pb.RegRequest, rsp *pb.RegResponse) error {
	// 注册用户，将数据写入到MySQL数据库
	err := model.RegistUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
	} else {
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	}
	return nil
}

