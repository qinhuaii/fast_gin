package flags

import (
	"flag"
)

type FlagOptions struct {
	File string
}

var Options FlagOptions

func Run() {
	// 启动时默认加载的配置文件
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	flag.Parse()
}
