package env

import (
	"os"
)

// Getenv 获取系统默认配置参数
// 如果没有返回默认值
func Getenv(env string, defaults string) string {
	if os.Getenv(env) == "" {
		return defaults
	}
	return os.Getenv(env)
}
