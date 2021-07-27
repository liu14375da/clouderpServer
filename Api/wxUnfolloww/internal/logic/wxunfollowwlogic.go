package logic

import (
	"ZeroProject/common/errorx"
	"context"
	"strings"

	"ZeroProject/Api/wxUnfolloww/internal/svc"
	"ZeroProject/Api/wxUnfolloww/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WxUnfollowwLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxUnfollowwLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxUnfollowwLogic {
	return WxUnfollowwLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxUnfollowwLogic) WxUnfolloww(req types.Request) (*types.Response, error) {
	if len(strings.TrimSpace(req.OpenId)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	if err := l.svcCtx.UnFollowModel.UpdataIsValid(req.OpenId);err != nil{
		return &types.Response{
			Message:"取消关注失败,"+err.Error(),
		}, nil
	}else {
		return &types.Response{
			Message:"取消关注成功",
		}, nil
	}
}
