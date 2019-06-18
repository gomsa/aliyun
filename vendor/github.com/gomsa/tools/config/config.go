package config

// Permission 权限
type Permission struct {
	Service     string `json:"service"`
	Method      string `json:"method"`
	Auth        bool   `json:"auth"`
	Policy      bool   `json:"policy"`
	Name        string `json:"display_name"`
	Description string `json:"description"`
}

// Config 配置结构
type Config struct {
	Service     string
	Version     string
	Permissions []Permission `json:"permissions"`
}
