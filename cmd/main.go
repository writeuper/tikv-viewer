package main

import (
	"log"

	"github.com/writeuper/tikv-viewer/internal/repository/tikv"

	"github.com/writeuper/tikv-viewer/config"
	"github.com/writeuper/tikv-viewer/pkg/logger"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 初始化日志
	logger.InitLogger(cfg.Log.Level)
	l := logger.GetLogger()

	// 初始化tikv客户端
	client, err := tikv.NewTiKVClient(cfg.TiKV.Addrs)
	if err != nil {
		l.Fatalf("failed to connect to tikv: %v", err)
		return
	}
	defer func(client *tikv.Client) {
		err := client.Close()
		if err != nil {
			l.Fatalf("failed to close tikv: %v", err)
		}
	}(client)

	// 初始化服务
	//tikvService := service.NewTiKVService(client)

	// 初始化api

}
