package k8s_info_router

import (
	"github.com/gin-gonic/gin"
	"k8s_controller/internal/handlers"
)

type K8sInfoHandler struct {
	handler *handlers.K8sInfoHandler
}

func NewK8sInfoRouter() *K8sInfoHandler {
	return &K8sInfoHandler{
		handler: handlers.NewK8sInfoHandler(),
	}
}
func (r *K8sInfoHandler) Register(router *gin.RouterGroup) {
	k8sInfoRoutes := router.Group("/k8sInfo")
	{
		k8sInfoRoutes.GET("/pods", r.handler.GetPostList)
		k8sInfoRoutes.GET("/podsStatus", r.handler.GetPodStatus)
	}
}
