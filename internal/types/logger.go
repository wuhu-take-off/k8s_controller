package types

import (
	"fmt"
	"k8s_controller/internal/global"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

// LoggerConfig 日志配置
type LoggerConfig struct {
	Logger struct {
		Level      string `yaml:"level"`       // 日志级别
		Filename   string `yaml:"filename"`    // 日志文件路径
		MaxSize    int    `yaml:"max-size"`    // 单个日志文件最大尺寸，单位MB
		MaxBackups int    `yaml:"max-backups"` // 最大保留日志文件数量
		MaxAge     int    `yaml:"max-age"`     // 日志文件保留天数
		Compress   bool   `yaml:"compress"`    // 是否压缩
	} `yaml:"Logger"`
}

// Load 加载日志配置
func (c *LoggerConfig) Load() error {
	// 读取配置文件
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(c.Logger.Filename), 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 设置日志级别
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(c.Logger.Level)); err != nil {
		return fmt.Errorf("解析日志级别失败: %v", err)
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	// 创建Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	global.Logger = logger

	return nil
}
