package user_router

import (
	"k8s_controller/internal/handlers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	handler *handlers.UserHandler
}

func NewUserRouter() *UserRouter {
	return &UserRouter{
		handler: handlers.NewUserHandler(),
	}
}

func (r *UserRouter) Register(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", r.handler.ListUsers)
		userRoutes.GET("/:id", r.handler.GetUser)
		userRoutes.POST("", r.handler.CreateUser)
		userRoutes.PUT("/:id", r.handler.UpdateUser)
		userRoutes.DELETE("/:id", r.handler.DeleteUser)
	}
}
