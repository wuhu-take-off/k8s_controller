package router

import (
	"github.com/gin-gonic/gin"
	"k8s_controller/internal/middleware"
	"k8s_controller/internal/router/k8s_info_router"
)

// SetupRouter 配置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Logger())

	// API 路由组
	api := r.Group("/api")
	{
		// 注册各个模块的路由
		registerRouters(api)
	}

	return r
}

// registerRouters 注册所有模块的路由
func registerRouters(router *gin.RouterGroup) {
	// 创建路由注册器列表
	routers := []Router{
		k8s_info_router.NewK8sInfoRouter(),
		// 在这里添加其他模块的路由注册器
	}

	// 注册所有路由
	for _, r := range routers {
		r.Register(router)
	}
}
