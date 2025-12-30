# Simple FTP Server

基于 Go 语言和 [ftpserverlib](https://github.com/fclairamb/ftpserverlib) 库实现的简易 FTP 服务器。本项目旨在演示如何使用该库快速构建一个支持 JSON 配置、用户认证及家目录隔离的定制化 FTP 服务。



## 核心依赖

本项目核心功能由 **[github.com/fclairamb/ftpserverlib](https://github.com/fclairamb/ftpserverlib)** 提供。
该库是一个功能强大且灵活的 Go 语言 FTP 服务器库，支持自定义驱动（Driver）来处理认证和文件系统操作。

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
      "home_dir": "/data/simple-ftp-server/storage/admin"
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
mkdir -p /data/simple-ftp-server/storage/admin

# 启动服务器 (默认加载同级目录下 config.json)
./simple-ftp-server

# 指定配置文件路径
./simple-ftp-server -config /data/simple-ftp-server/config.json
```

## 跨平台编译方法

使用 Go 语言的交叉编译功能，可以轻松生成不同平台的二进制文件。

### 编译 Linux 版本 (amd64)
适用于常见的 Linux 服务器 (CentOS, Ubuntu 等)。
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-ftp-server-linux-amd64 main.go
```

### 编译 Windows 版本 (amd64)
生成 `.exe` 文件。
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o simple-ftp-server-windows-amd64.exe main.go
```

### 编译 macOS 版本 (amd64 / Intel)
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o simple-ftp-server-mac-intel main.go
```

### 编译 macOS 版本 (arm64 / Apple Silicon)
适用于 M1/M2/M3 芯片。
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o simple-ftp-server-mac-m1 main.go
```

## Systemd 自动开机启动 (Linux)

本项目提供了 `simple-ftp-server.service` 模板文件，用于配置 Systemd 服务。

### 1. 部署文件
假设你将程序部署在 `/data/simple-ftp-server` 目录：
```bash
# 创建目录
mkdir -p /data/simple-ftp-server/

# 复制文件
cp simple-ftp-server-linux-amd64 /data/simple-ftp-server/simple-ftp-server
cp config.json /data/simple-ftp-server/
cp simple-ftp-server.service /etc/systemd/system/
```

### 2. 修改配置
编辑 `/etc/systemd/system/simple-ftp-server.service`，确保路径正确：
```ini
[Service]
WorkingDirectory=/data/simple-ftp-server
ExecStart=/data/simple-ftp-server/simple-ftp-server -config /data/simple-ftp-server/config.json
```

### 3. 启动服务
```bash
# 重载配置
systemctl daemon-reload

# 启动服务
systemctl start simple-ftp-server

#在此设置开机自启
systemctl enable simple-ftp-server

# 查看状态
systemctl status simple-ftp-server
```


