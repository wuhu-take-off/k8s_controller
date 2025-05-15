package types

import (
	"encoding/base64"
	"fmt"
	"k8s_controller/internal/global"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// K8sConfig K8s配置
type K8sConfig struct {
	// K8sConfig K8s配置信息
	K8sConfig struct {
		Cluster struct {
			Certificate string `yaml:"Certificate"`
			Server      string `yaml:"Server"`
			ClusterName string `yaml:"ClusterName"`
		} `yaml:"Cluster"`
		Context struct {
			Cluster     string `yaml:"Cluster"`
			User        string `yaml:"User"`
			ContextName string `yaml:"ContextName"`
		} `yaml:"Context"`
		CurrentContext string `yaml:"CurrentContext"`
		Kind           string `yaml:"Kind"`
		Users          struct {
			UsersName             string `yaml:"UsersName"`
			ClientCertificateData string `yaml:"ClientCertificateData"`
			ClientKeyData         string `yaml:"ClientKeyData"`
		} `yaml:"Users"`
	} `yaml:"K8sConfig"`
	// K8sMaster K8s主节点配置
	K8sMaster struct {
		IPAddress     string `yaml:"ip-address"`
		Port          string `yaml:"port"`
		Username      string `yaml:"username"`
		Password      string `yaml:"password"`
		CephSC        string `yaml:"ceph-sc"`
		FreeDataSize  int    `yaml:"free-data-size"`
		FreeDataMount string `yaml:"free-data-mount"`
	} `yaml:"K8sMaster"`
}

// Load 实现配置加载
func (c K8sConfig) Load() error {
	// 验证K8s配置
	if c.K8sConfig.Cluster.Server == "" {
		return fmt.Errorf("K8s服务器地址不能为空")
	}

	// 验证主节点配置
	if c.K8sMaster.IPAddress == "" {
		return fmt.Errorf("K8s主节点IP地址不能为空")
	}
	if c.K8sMaster.Port == "" {
		return fmt.Errorf("K8s主节点端口不能为空")
	}
	if c.K8sMaster.Username == "" {
		return fmt.Errorf("K8s主节点用户名不能为空")
	}
	if c.K8sMaster.Password == "" {
		return fmt.Errorf("K8s主节点密码不能为空")
	}

	// 解码证书数据
	clusterCert, err := base64.StdEncoding.DecodeString(c.K8sConfig.Cluster.Certificate)
	if err != nil {
		return fmt.Errorf("解码集群证书失败: %v", err)
	}

	clientCert, err := base64.StdEncoding.DecodeString(c.K8sConfig.Users.ClientCertificateData)
	if err != nil {
		return fmt.Errorf("解码客户端证书失败: %v", err)
	}

	clientKey, err := base64.StdEncoding.DecodeString(c.K8sConfig.Users.ClientKeyData)
	if err != nil {
		return fmt.Errorf("解码客户端密钥失败: %v", err)
	}

	// 创建K8s配置
	config := api.NewConfig()

	// 设置集群信息
	config.Clusters[c.K8sConfig.Cluster.ClusterName] = &api.Cluster{
		Server:                   c.K8sConfig.Cluster.Server,
		CertificateAuthorityData: clusterCert,
	}

	// 设置用户信息
	config.AuthInfos[c.K8sConfig.Users.UsersName] = &api.AuthInfo{
		ClientCertificateData: clientCert,
		ClientKeyData:         clientKey,
	}

	// 设置上下文信息
	config.Contexts[c.K8sConfig.Context.ContextName] = &api.Context{
		Cluster:  c.K8sConfig.Context.Cluster,
		AuthInfo: c.K8sConfig.Users.UsersName,
	}

	// 设置当前上下文
	config.CurrentContext = c.K8sConfig.CurrentContext

	// 创建客户端配置
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{
		AuthInfo: api.AuthInfo{
			ClientCertificateData: clientCert,
			ClientKeyData:         clientKey,
		},
		ClusterInfo: api.Cluster{
			Server:                   c.K8sConfig.Cluster.Server,
			CertificateAuthorityData: clusterCert,
		},
		Context: api.Context{
			Cluster:  c.K8sConfig.Context.Cluster,
			AuthInfo: c.K8sConfig.Users.UsersName,
		},
		CurrentContext: c.K8sConfig.CurrentContext,
	})
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("创建K8s配置失败: %v", err)
	}

	// 设置QPS和Burst
	restConfig.QPS = 100
	restConfig.Burst = 200

	// 保存REST配置
	global.K8sRestConfig = restConfig

	// 创建K8s客户端
	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return fmt.Errorf("创建K8s客户端失败: %v", err)
	}

	// 保存客户端
	global.K8sClient = client
	return nil
}
