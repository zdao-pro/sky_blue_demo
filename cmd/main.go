package main

import (
	server "sky_blue_demo/internal/server/http"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
	"sky_blue_demo/internal/dao"
)

func main() {
	log.Init(nil)
	log.Info("start the server....")
	//配置中心初始化,设置apollo作为驱动
	err := peach.Init(peach.PeachDriverApollo, []string{"zdao_backend.sky_blue", "zdao_backend.common"})
	if nil != err {
		panic(err)
	}
    dao.Init()
	httpServer := server.NewServer()
	//开始运行http server
	httpServer.Run(":8080")
}
