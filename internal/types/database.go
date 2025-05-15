package types

import "fmt"

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `yaml:"driver"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

// Load 实现配置加载
func (c DatabaseConfig) Load() error {
	fmt.Println("加载数据库配置信息")
	// 验证驱动类型
	if c.Driver == "" {
		return fmt.Errorf("数据库驱动类型不能为空")
	}

	// 验证主机地址
	if c.Host == "" {
		return fmt.Errorf("数据库主机地址不能为空")
	}

	// 验证端口号
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("无效的数据库端口号: %d", c.Port)
	}

	// 验证用户名
	if c.Username == "" {
		return fmt.Errorf("数据库用户名不能为空")
	}

	// 验证数据库名
	if c.DBName == "" {
		return fmt.Errorf("数据库名称不能为空")
	}

	// 验证连接池配置
	if c.MaxIdleConns <= 0 {
		return fmt.Errorf("无效的最大空闲连接数: %d", c.MaxIdleConns)
	}
	if c.MaxOpenConns <= 0 {
		return fmt.Errorf("无效的最大打开连接数: %d", c.MaxOpenConns)
	}
	if c.ConnMaxLifetime <= 0 {
		return fmt.Errorf("无效的连接最大生命周期: %d", c.ConnMaxLifetime)
	}

	return nil
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.DBName)
}
