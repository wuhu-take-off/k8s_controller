package main

import (
	"fmt"
	"k8s_controller/config"
	"k8s_controller/internal/router"
	"log"
)

func main() {
	// 初始化配置
	if err := config.Init(""); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
