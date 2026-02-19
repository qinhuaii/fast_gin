package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
)

func main() {
	core.InitLogger() // 日志记录
	flags.Parse()
	global.Config = core.ReadConfig() // 读取yaml文件里的配置
	global.DB = core.InitGorm()       // 连接数据库
	global.Redis = core.InitRedis()   // Redis连接
	flags.Run()                       // -db 表结构迁移 | -v 打印版本信息 | -f 修改默认读取的配置文件
}
