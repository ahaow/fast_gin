package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
)

func main() {
	flags.Parse()

	// 1. 读取配置
	global.Config = core.InitConfig(flags.Options.Env)

	// 2. 初始化日志
	log, err := core.NewLogger("development", "logs", "[myApp] ")
	if err != nil {
		panic(err)
	}
	global.Log = log

	defer log.Close() // 确保程序结束时同步日志

	// 3. 连接数据库
	global.DB = core.InitGorm(global.Config)

	// 4. 连接redis
	// global.Redis = core.InitRedis(global.Config)

	// 5. gin
	routers.Run()

	flags.Run()

	// fmt.Println(global.Config.App.Name)
}
