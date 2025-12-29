# Simple FTP Server

基于 Go 语言和 `github.com/fclairamb/ftpserverlib` 库实现的简易 FTP 服务器。支持 JSON 配置、用户认证及家目录隔离。

## 功能特性

*   **轻量级**: 基于 Go 语言开发，单文件部署。
*   **配置灵活**: 使用 `config.json` 文件进行配置。
*   **用户认证**: 支持基于 JSON 文件的用户名/密码认证。
*   **目录隔离**: 每个用户只能访问其指定的家目录 (Chroot)。
*   **网络模式**: 支持主动模式和被动模式 (Passive Mode)，可自定义被动模式端口范围和公网 IP。
*   **全中文注释**: 代码中包含详细的中文注释，便于学习和二次开发。

## 快速开始

### 1. 编译
```bash
go mod tidy
go build -o ftp-server main.go
```

### 2. 配置
在可执行文件同级目录下创建 `config.json` 文件：

```json
{
  "server": {
    "listen_addr": ":2121",
    "public_ip": "127.0.0.1",
    "passive_port_low": 30000,
    "passive_port_high": 30010
  },
  "users": [
    {
      "username": "admin",
      "password": "password123",
      "home_dir": "/tmp/ftp/admin"
    }
  ]
}
```

**配置项说明**:
*   `server.listen_addr`: 服务器监听地址和端口。
*   `server.public_ip`: (可选) 被动模式下告知客户端的公网 IP。如果在内网或本地测试，可填 `127.0.0.1`。
*   `server.passive_port_low` / `high`: 被动模式数据传输使用的端口范围。确保防火墙放行这些端口。
*   `users`: 用户列表。
    *   `username`: 登录用户名。
    *   `password`: 登录密码。
    *   `home_dir`: 该用户的根目录绝对路径。**注意：该目录必须存在，服务器启动前请确保已创建。**

### 3. 运行
```bash
# 确保家目录存在
mkdir -p /tmp/ftp/admin

# 启动服务器
./ftp-server
```

## 跨平台编译方法

使用 Go 语言的交叉编译功能，可以轻松生成不同平台的二进制文件。

### 编译 Linux 版本 (amd64)
适用于常见的 Linux 服务器 (CentOS, Ubuntu 等)。
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ftp-server-linux main.go
```

### 编译 Windows 版本 (amd64)
生成 `.exe` 文件。
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ftp-server-windows.exe main.go
```

### 编译 macOS 版本 (amd64 / Intel)
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ftp-server-mac-intel main.go
```

### 编译 macOS 版本 (arm64 / Apple Silicon)
适用于 M1/M2/M3 芯片。
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ftp-server-mac-m1 main.go
```
