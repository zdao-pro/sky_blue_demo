package main

import (
	server "sky_blue_demo/internal/server/http"

	"sky_blue_demo/internal/dao"

	"context"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/net/http/request"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func main() {
	// 初始化日志模块
	log.Init(nil)
	//建议使用Infoc,Warnc....,带上参数context,
	//对于gin来说,gin.Context初始时，c.Context会存储trace对象，方便日志打印trace_id
	log.Infoc(context.Background(), "start the server....")

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
