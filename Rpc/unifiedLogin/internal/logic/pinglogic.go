package logic

import (
	pb "ZeroProject/Rpc/unifiedLogin/pb"
	"ZeroProject/common/global"
	"ZeroProject/common/jwt"
	"ZeroProject/common/tool"
	"ZeroProject/model/sql"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"ZeroProject/Rpc/unifiedLogin/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func (l LoginTokenLogic) GetLoginToken(req *pb.Request) (*pb.Response, error) {
	md5str2 := tool.UniversalEncryption()
	userInfo, err := l.svcCtx.UserModel.FindUnified(req.Username)
	switch err {
	case nil:
	case sql.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}
	if userInfo.UserPassword != req.Password && req.Password != strings.ToUpper(md5str2) {
		return nil, errors.New("用户密码不正确")
	}
	fmt.Println(global.JwtAuth)
	//生成token
	now := time.Now().Unix()
	accessExpire := global.JwtAuth.AccessExpire
	jwtToken, err := jwt.GetJwtToken(
		now,
		global.JwtAuth.AccessExpire,
		global.JwtAuth.AccessSecret,
		userInfo.UserName,
		userInfo.CompanyId,
		userInfo.StaffId,
		userInfo.UserId,
	)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Token:  jwtToken,
		Expire: now + accessExpire,
	}, nil
}

func NewLoginToken(c context.Context, svc *svc.ServiceContext) *LoginTokenLogic {
	return &LoginTokenLogic{
		ctx:    c,
		svcCtx: svc,
		Logger: logx.WithContext(c),
	}
}
