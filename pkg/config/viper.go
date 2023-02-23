package config

import (
	"fmt"
	"github.com/spf13/viper"
)

//https://github.com/spf13/viper#reading-config-files

func InitConfig(filepath string) {
	viper.SetConfigName("config") // 设置配置文件名称 name of config file (without extension)
	viper.SetConfigType("yaml")   // 设置配置文件类型 也就是后缀 REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(filepath) // 设置 配置文件路径 path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return
}
