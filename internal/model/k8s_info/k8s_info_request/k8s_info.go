package k8s_info_request

type PodsListRequest struct {
}
type PodsStatusRequest struct {
	Namespace string   `json:"namespace"` // 命名空间
	NodeName  []string `json:"nodeName"`  // 节点名称
}
