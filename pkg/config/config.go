package config

import (
	"encoding/json"
	"os"
)

// Config 定义了整个应用程序的配置结构
type Config struct {
	Server ServerConfig `json:"server"` // 服务器配置
	Users  []User       `json:"users"`  // 用户列表
}

// ServerConfig 定义了 FTP 服务器的网络配置
type ServerConfig struct {
	ListenAddr      string `json:"listen_addr"`      // 监听地址，例如 ":2121"
	PublicIP        string `json:"public_ip"`        // 被动模式使用的公网 IP (可选)
	PassivePortLow  int    `json:"passive_port_low"` // 被动模式端口范围起始
	PassivePortHigh int    `json:"passive_port_high"`// 被动模式端口范围结束
}

// User 定义了 FTP 用户及其权限
type User struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码 (明文，实际生产中应使用哈希)
	HomeDir  string `json:"home_dir"` // 用户家目录的绝对路径
}

// LoadConfig 从指定的 JSON 文件加载配置
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
