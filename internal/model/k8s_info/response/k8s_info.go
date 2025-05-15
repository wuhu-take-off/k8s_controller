package k8s_info_response

import (
	"k8s_controller/pkg/utils"
)

type PodsListResponse struct {
	NodeName  string            `json:"node_name"`
	Status    string            `json:"status"` // pod状态
	PodIp     string            `json:"pod_ip"`
	CreatedAt utils.LocalTime   `json:"created_at"`
	Labels    map[string]string `json:"labels"`
}
