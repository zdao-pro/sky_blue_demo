package main

import (
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
}
