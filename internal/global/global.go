package global

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	// K8sConfig K8s配置
	K8sClient *kubernetes.Clientset
	// K8sRestConfig K8s REST配置
	K8sRestConfig *rest.Config
	GVA_LOG       *zap.Logger
	Logger        *zap.Logger
)
