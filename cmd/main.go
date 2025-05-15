package main

import (
	"github.com/gin-gonic/gin"
	"github.com/writeuper/tikv-viewer/internal/service"
	"log"
	"net/http"

	"github.com/writeuper/tikv-viewer/internal/repository/tikv"

	"github.com/writeuper/tikv-viewer/api/v1"
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
	tikvService := service.NewTiKVService(client)

	// 初始化api
	tikvAPI := v1.NewTikvAPI(tikvService)

	// 创建gin
	r := gin.Default()

	// 注册API路由
	apiGroup := r.Group("/api/v1/tikv")
	{
		apiGroup.GET("/:key", tikvAPI.GetValueHandler)
		apiGroup.POST("/", tikvAPI.SetValueHandler)
		apiGroup.DELETE("/:key", tikvAPI.DeleteValueHandler)
		apiGroup.POST("/scan", tikvAPI.ScanHandler)
	}

	// 静态文件服务
	r.Static("/static", "./web/dist")

	// 根路径重定向到静态首页
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/index.html")
	})

	// 启动服务器
	l.Infof("Server started on %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil && err != http.ErrServerClosed {
		l.Fatalf("Server failed to start: %v", err)
	}
}
