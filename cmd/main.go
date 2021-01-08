package main

import (
	server "sky_blue_demo/internal/server/http"

	"sky_blue_demo/internal/dao"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/request"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func main() {
	// 初始化日志模块
	log.Init(nil)
	log.Info("start the server....")

	// 配置中心初始化,设置apollo作为驱动
	err := peach.Init(peach.PeachDriverApollo, []string{"zdao_backend.sky_blue", "zdao_backend.common"})
	if nil != err {
		panic(err)
	}

	// 初始化dao
	dao.Init()

	// 初始化http request
	upstreamStr, _ := peach.Get("upstream.yaml").String()
	request.InitUpstream(upstreamStr)

	// 开始运行http server
	httpServer := server.NewServer()
	httpServer.Run(":8080")
}
