package types

import (
	"fmt"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         int    `yaml:"port"`
	Mode         string `yaml:"mode"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// Load 实现配置加载
func (c ServerConfig) Load() error {
	fmt.Println("加载server配置信息")
	// 验证端口号
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("无效的端口号: %d", c.Port)
	}

	// 验证模式
	if c.Mode != "debug" && c.Mode != "release" {
		return fmt.Errorf("无效的运行模式: %s，必须是 debug 或 release", c.Mode)
	}

	// 验证超时时间
	if c.ReadTimeout <= 0 {
		return fmt.Errorf("无效的读取超时时间: %d", c.ReadTimeout)
	}
	if c.WriteTimeout <= 0 {
		return fmt.Errorf("无效的写入超时时间: %d", c.WriteTimeout)
	}

	return nil
}
