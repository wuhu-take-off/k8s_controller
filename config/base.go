package config

import (
	"fmt"
	"k8s_controller/internal/types"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v2"
)

var (
	// GlobalConfig 全局配置
	GlobalConfig types.Config
)

const (
	DefaultConfigFile = "config/config.yaml"
)

// Init 初始化配置
func Init(configPath string) error {
	// 如果未指定配置文件路径，使用默认路径
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 加载配置
	if err := loadConfig(&GlobalConfig); err != nil {
		return fmt.Errorf("加载配置失败: %v", err)
	}

	return nil
}

// loadConfig 使用反射加载配置
func loadConfig(config interface{}) error {
	// 获取配置对象的反射值
	val := reflect.ValueOf(config)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("配置对象必须是结构体指针")
	}

	// 遍历结构体字段
	val = val.Elem()
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 检查字段是否实现了 ConfigLoader 接口
		if field.CanAddr() {
			fieldAddr := field.Addr()
			if loader, ok := fieldAddr.Interface().(types.ConfigLoader); ok {
				if err := loader.Load(); err != nil {
					return fmt.Errorf("加载%s配置失败: %v", fieldType.Name, err)
				}
			}
		}
	}

	return nil
}

// ensureLogDir 确保日志目录存在
func ensureLogDir(filename string) error {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
