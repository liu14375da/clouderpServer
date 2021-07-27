// Code generated by goctl. DO NOT EDIT!
// Source: unifiedLogin.proto

package server

import (
	"ZeroProject/Rpc/unifiedLogin/internal/logic"
	"ZeroProject/Rpc/unifiedLogin/internal/svc"
	pb "ZeroProject/Rpc/unifiedLogin/pb"
	"context"
)

type UnifiedLoginServer struct {
	svcCtx *svc.ServiceContext
}

func (u UnifiedLoginServer) LoginToken(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	l := logic.NewLoginToken(ctx, u.svcCtx)
	return l.GetLoginToken(request)
}

func NewUnifiedLoginServer(svcCtx *svc.ServiceContext) *UnifiedLoginServer {
	return &UnifiedLoginServer{
		svcCtx: svcCtx,
	}
}