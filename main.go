package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
)

func main() {
	core.InitLogger() // 日志记录
	flags.Parse()
	global.Config = core.ReadConfig() // 读取yaml文件里的配置
	global.DB = core.InitGorm()       // 连接数据库
	global.Redis = core.InitRedis()   // Redis连接
	// -db 表结构迁移 | -v 打印版本信息 | -f 修改默认读取的配置文件 | -m 选择命令行模式 | -t 选择命令行类型
	flags.Run()
	routers.Run()
}
