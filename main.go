package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"simple-ftp-server/pkg/config"
	"simple-ftp-server/pkg/driver"

	ftpserver "github.com/fclairamb/ftpserverlib"
)

func main() {
	// 0. 解析参数
	configPath := flag.String("config", "config.json", "配置文件路径")
	flag.Parse()

	// 1. 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}

	// 2. 初始化驱动
	d := driver.NewMainDriver(cfg)

	// 3. 创建 FTP 服务器
	server := ftpserver.NewFtpServer(d)

	// 4. 启动服务器
	go func() {
		log.Printf("FTP 服务器已启动，监听地址: %s", cfg.Server.ListenAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("无法启动服务器: %v", err)
		}
	}()

	// 5. 等待中断信号以优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")
	if err := server.Stop(); err != nil {
		log.Printf("关闭服务器时出错: %v", err)
	}
	log.Println("服务器已关闭")
}
