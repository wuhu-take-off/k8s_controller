package router

import "github.com/gin-gonic/gin"

// Router 路由注册接口
type Router interface {
	Register(r *gin.RouterGroup)
}

// RegisterRouter 注册路由的函数类型
type RegisterRouter func(r *gin.RouterGroup)
