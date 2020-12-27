package main

import (
	server "sky_blue_demo/server/http"

	"github.com/zdao-pro/sky_blue/pkg/log"
	"github.com/zdao-pro/sky_blue/pkg/peach"
	_ "github.com/zdao-pro/sky_blue/pkg/peach/apollo"
)

func main() {
	log.Init(nil)
	err := peach.Init(peach.PeachDriverApollo, "zdao_backend.sky_blue")
	if nil != err {
		panic(err)
	}

	httpServer := server.NewServer()
	httpServer.Run()
}
