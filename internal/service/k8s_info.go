package service

import (
	"context"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_controller/internal/global"
	"k8s_controller/internal/model/k8s_info/k8s_info_request"
	k8s_info_response "k8s_controller/internal/model/k8s_info/response"
	"k8s_controller/pkg/utils"
)

type K8sInfoService struct {
}

func NewK8sInfoService() *K8sInfoService {
	return &K8sInfoService{}
}

const (
	namespace = "person-test"
)

func (s K8sInfoService) GetPodList(req k8s_info_request.PodsListRequest) ([]*k8s_info_response.PodsListResponse, int64, error) {
	pods, err := global.K8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		global.Logger.Error("获取pods信息失败", zap.Error(err))
		return nil, 0, utils.NewErrWithMessage("pods信息获取失败")
	}
	res := make([]*k8s_info_response.PodsListResponse, 0)
	for i := range pods.Items {
		res = append(res, &k8s_info_response.PodsListResponse{
			NodeName:  pods.Items[i].Name,
			Status:    string(pods.Items[i].Status.Phase),
			PodIp:     pods.Items[i].Status.PodIP,
			CreatedAt: utils.LocalTime(pods.Items[i].CreationTimestamp.Time),
			Labels:    pods.Items[i].Labels,
		})
	}
	return res, int64(len(pods.Items)), nil
}

func (s K8sInfoService) GetPodStatus(req k8s_info_request.PodsStatusRequest) error {
	return nil
}
