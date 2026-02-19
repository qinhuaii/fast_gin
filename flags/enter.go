package flags

import (
	"fast_gin/global"
	"flag"
	"fmt"
	"os"
)

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
}

var Options FlagOptions

func Parse() {
	// 启动时默认加载的配置文件
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	// 加载版本信息
	flag.BoolVar(&Options.Version, "v", false, "打印当前的版本")
	// 迁移表结构
	flag.BoolVar(&Options.DB, "db", false, "迁移表结构")
	flag.Parse()
}

func Run() {
	if Options.DB {
		MigrateDB()
		os.Exit(0)
	}
	if Options.Version {
		fmt.Println("当前后端版本：", global.Version)
		os.Exit(0)
	}
}
