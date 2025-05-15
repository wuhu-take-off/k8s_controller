package handlers

import (
	"k8s_controller/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	// TODO: 实现获取用户列表
	c.JSON(200, gin.H{
		"message": "list users",
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// TODO: 实现获取单个用户
	c.JSON(200, gin.H{
		"message": "get user",
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// TODO: 实现创建用户
	c.JSON(200, gin.H{
		"message": "create user",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// TODO: 实现更新用户
	c.JSON(200, gin.H{
		"message": "update user",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// TODO: 实现删除用户
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}
