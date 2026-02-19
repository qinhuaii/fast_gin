package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
)

func main() {
	core.InitLogger()
	flags.Run()
	global.Config = core.ReadConfig()

}
