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
	Menu    string // 菜单
	Type    string // 类型 create list remove
}

var Options FlagOptions

func Parse() {
	// 启动时默认加载的配置文件
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	// 执行命令操作
	flag.StringVar(&Options.Menu, "m", "", "菜单 user")
	// 命令操作的类型
	flag.StringVar(&Options.Type, "t", "", "类型 create list")
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
	// 用户相关的操作
	if Options.Menu == "user" {
		var user User
		switch Options.Type {
		case "create":
			user.Create()
		case "list":
			user.List()
		}
		os.Exit(0)
	}
}
