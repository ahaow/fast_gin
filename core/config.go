package core

import (
	"fast_gin/config"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig(envString string) *config.Config {

	env := strings.ToLower(envString)

	// 构建 config 文件名
	configName := "config"
	if env == "dev" {
		configName = "config_dev"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置失败 [%s.yaml]: %v", configName, err)
	}

	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("配置解析失败: %v", err)
	}

	fmt.Println("配置加载成功:", viper.ConfigFileUsed())
	return cfg
}
