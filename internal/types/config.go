package types

// ConfigLoader 配置加载接口
type ConfigLoader interface {
	Load() error
}

// Config 总配置结构体
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      LoggerConfig   `yaml:"log"`
	K8s      K8sConfig      `yaml:"k8s"`
}
