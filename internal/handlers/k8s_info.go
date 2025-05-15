package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"k8s_controller/internal/model/common/common_response"
	"k8s_controller/internal/model/k8s_info/k8s_info_request"
	"k8s_controller/internal/model/response"
	"k8s_controller/internal/service"
	"k8s_controller/pkg/utils"
)

type K8sInfoHandler struct {
	K8sInfoService *service.K8sInfoService
}

func NewK8sInfoHandler() *K8sInfoHandler {
	return &K8sInfoHandler{
		K8sInfoService: service.NewK8sInfoService(),
	}
}
func (h *K8sInfoHandler) GetPostList(c *gin.Context) {
	var req k8s_info_request.PodsListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewResponseWithErrorString("参数解析失败", c)
		return
	}
	pods, total, err := h.K8sInfoService.GetPodList(req)
	if err != nil {
		var uErr *utils.Err
		if errors.As(err, &uErr) {
			response.NewResponseWithErrorString(uErr.Error(), c)
			return
		} else {
			response.NewResponseWithErrorString("获取失败", c)
		}
		return
	}
	response.NewResponseWithData(common_response.PageResponse{
		Total: total,
		List:  pods,
	}, c)
}

func (h *K8sInfoHandler) GetPodStatus(c *gin.Context) {
	var req k8s_info_request.PodsStatusRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewResponseWithErrorString("参数解析失败", c)
		return
	}

	err := h.K8sInfoService.GetPodStatus(req)
	if err != nil {
		var uErr *utils.Err
		if errors.As(err, &uErr) {
			response.NewResponseWithErrorString(uErr.Error(), c)
			return
		} else {
			response.NewResponseWithErrorString("获取失败", c)
		}
		return
	}
	response.NewResponseWithData("", c)

}
