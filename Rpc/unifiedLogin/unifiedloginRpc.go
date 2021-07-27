package main

import (
	"ZeroProject/Rpc/unifiedLogin/internal/server"
	"ZeroProject/Rpc/unifiedLogin/internal/svc"
	pb "ZeroProject/Rpc/unifiedLogin/pb"
	nacos "ZeroProject/nacos/server"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/zrpc"

	"google.golang.org/grpc"
)

/*
	统一登录
*/
func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	c = svc.RpcServer(c)

	ctx := svc.NewServiceContext(c)
	srv := server.NewUnifiedLoginServer(ctx)

	s := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		pb.RegisterUnifiedLoginServer(grpcServer, srv)
	})

	defer s.Stop()

	nacos.RegisterService("UnifiedLogin", nil)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
