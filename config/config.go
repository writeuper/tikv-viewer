package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	TiKV struct {
		Addrs []string `mapstructure:"addrs"`
	} `mapstructure:"tikv"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}

func LoadConfig() (Config, error) {
	config := Config{}

	// viper进行解析
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("tikv.addrs", []string{"127.0.0.1:2379"})
	viper.SetDefault("log.level", "info")

	// 读取环境变量
	viper.AutomaticEnv()

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalf("Config file not found: %v", err)
		} else {
			return config, fmt.Errorf("Error reading config file: %v", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("Error unmarshaling config: %v", err)
	}

	// 应用环境变量覆盖
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}

	if addrs := os.Getenv("TIKV_ADDRS"); addrs != "" {
		config.TiKV.Addrs = []string{addrs}
	}

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Log.Level = level
	}
	return config, nil
}
