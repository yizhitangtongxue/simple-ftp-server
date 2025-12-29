package driver

import (
	"crypto/tls"
	"errors"
	"log"

	ftpserver "github.com/fclairamb/ftpserverlib"
	"github.com/spf13/afero"

	"simple-ftp-server/pkg/config"
)

// MainDriver 实现了 ftpserverlib.MainDriver 接口
// 负责服务器配置、用户认证和客户端连接管理
type MainDriver struct {
	config *config.Config
}

// NewMainDriver 创建一个新的 MainDriver 实例
func NewMainDriver(cfg *config.Config) *MainDriver {
	return &MainDriver{
		config: cfg,
	}
}

// GetSettings 返回 FTP 服务器的配置
func (d *MainDriver) GetSettings() (*ftpserver.Settings, error) {
	var portRange *ftpserver.PortRange
	if d.config.Server.PassivePortLow > 0 && d.config.Server.PassivePortHigh > 0 {
		portRange = &ftpserver.PortRange{
			Start: d.config.Server.PassivePortLow,
			End:   d.config.Server.PassivePortHigh,
		}
	}

	return &ftpserver.Settings{
		ListenAddr:               d.config.Server.ListenAddr,
		PublicHost:               d.config.Server.PublicIP,
		PassiveTransferPortRange: portRange,
		// 这里可以添加更多配置，例如 TLS
	}, nil
}

// ClientConnected 当有新客户端连接时被调用
func (d *MainDriver) ClientConnected(cc ftpserver.ClientContext) (string, error) {
	log.Printf("客户端连接: %s", cc.RemoteAddr())
	return "欢迎使用 Go FTP 服务器", nil
}

// ClientDisconnected 当客户端断开连接时被调用
func (d *MainDriver) ClientDisconnected(cc ftpserver.ClientContext) {
	log.Printf("客户端断开: %s", cc.RemoteAddr())
}

// AuthUser 验证用户身份
func (d *MainDriver) AuthUser(cc ftpserver.ClientContext, user, pass string) (ftpserver.ClientDriver, error) {
	for _, u := range d.config.Users {
		if u.Username == user && u.Password == pass {
			log.Printf("用户登录成功: %s", user)
			// 使用 afero.NewBasePathFs 将用户文件系统限制在其家目录下
			fs := afero.NewBasePathFs(afero.NewOsFs(), u.HomeDir)
			return &ClientDriver{Fs: fs}, nil
		}
	}
	log.Printf("用户登录失败: %s", user)
	return nil, errors.New("用户名或密码错误")
}

// GetTLSConfig 返回 TLS 配置 (目前不支持)
func (d *MainDriver) GetTLSConfig() (*tls.Config, error) {
	return nil, errors.New("TLS not configured")
}

// ClientDriver 实现了 ftpserverlib.ClientDriver 接口
// 并直接嵌入 afero.Fs 以处理底层文件系统操作
type ClientDriver struct {
	afero.Fs
}

// Allocate 预分配空间 (可选实现，这里留空)
func (d *ClientDriver) Allocate(space int) error {
	return nil
}
