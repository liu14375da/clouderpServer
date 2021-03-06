package main

import (
	"ZeroProject/Api/wxUnfolloww/internal/handler"
	"ZeroProject/Api/wxUnfolloww/internal/svc"
	"ZeroProject/common/middleware"
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/rest"
)

/*
	取消关注
*/
func main() {
	flag.Parse()

	var c svc.Config
	c = svc.Config(svc.ClientConfig(c))

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf,
		rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()
	//跨域
	server.Use(middleware.NewCorsMiddleware().Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
