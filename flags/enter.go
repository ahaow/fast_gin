package flags

import (
	"flag"
	"fmt"
	"os"
)

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
	Env     string
	// Menu    string // 菜单
	// Type    string // 类型 create
}

var Options FlagOptions

func Parse() {
	flag.StringVar(&Options.File, "f", "config.yaml", "配置文件路径")
	// flag.StringVar(&Options.Menu, "m", "", "菜单 user")
	// flag.StringVar(&Options.Type, "t", "", "类型 create list")
	flag.BoolVar(&Options.Version, "v", true, "打印当前版本")
	flag.BoolVar(&Options.DB, "db", true, "迁移表结构")
	flag.StringVar(&Options.Env, "env", "prod", "环境变量：dev 或 prod") // ✅ 新增 env
	flag.Parse()
}

func Run() {
	if Options.DB {
		MigrateDB()
		fmt.Println("表结构迁移")
		os.Exit(0)
	}
	if Options.Version {
		fmt.Println("当前后端版本号")
		os.Exit(0)
	}

	// if Options.Menu == "user" {
	// 	var user User
	// 	switch Options.Type {
	// 	case "create":
	// 		user.Create()
	// 	case "list":
	// 		user.List()
	// 	}
	// 	os.Exit(0)
	// }

}
