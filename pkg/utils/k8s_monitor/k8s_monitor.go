package k8s_monitor

import (
	"k8s.io/apimachinery/pkg/watch"
	"sync"
)

type K8sMonitor struct {
	watch watch.Interface
}

var (
	once       sync.Once
	k8sMonitor *K8sMonitor
)

func NewK8sMonitor() *K8sMonitor {
	once.Do(func() {
		k8sMonitor = &K8sMonitor{}
		//k8sMonitor.watch = global.K8sClient.CoreV1().Events()
	})
	return nil
}
