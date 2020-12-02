package nodeconfigs

import "github.com/iwind/TeaGo/rands"

// 默认的TOA配置
func DefaultTOAConfig() *TOAConfig {
	return &TOAConfig{
		IsOn:       false,
		Debug:      false,
		OptionType: 0xfe,
		MinQueueId: 100,
		MaxQueueId: 109,
		AutoSetup:  true,
	}
}

// TOA相关配置
type TOAConfig struct {
	IsOn         bool     `yaml:"isOn" json:"isOn"` // 是否启用
	Debug        bool     `yaml:"debug" json:"debug"`
	OptionType   uint8    `yaml:"optionType" json:"optionType"`
	MinQueueId   uint8    `yaml:"minQueueId" json:"minQueueId"`
	MaxQueueId   uint8    `yaml:"maxQueueId" json:"maxQueueId"`
	AutoSetup    bool     `yaml:"autoSetup" json:"autoSetup"`
	MinLocalPort uint16   `yaml:"minLocalPort" json:"minLocalPort"` // 本地可使用的最小端口 TODO
	MaxLocalPort uint16   `yaml:"maxLocalPort" json:"maxLocalPort"` // 本地可使用的最大端口 TODO
	SockPath     string   `yaml:"sockPath" json:"sockPath"`         // Sock文件路径 TODO
	ByPassPorts  []uint16 `yaml:"byPassPorts" json:"byPassPorts"`   // 忽略的端口 TODO

	minLocalPort int
	maxLocalPort int
}

func (this *TOAConfig) Init() error {
	// LocalPort
	minPort := this.MinLocalPort
	maxPort := this.MaxLocalPort
	if minPort == 0 {
		minPort = 1025
	}
	if maxPort == 0 {
		maxPort = 65534
	}
	if minPort > maxPort {
		minPort, maxPort = maxPort, minPort
	}
	this.minLocalPort = int(minPort)
	this.maxLocalPort = int(maxPort)

	return nil
}

// Sock路径
func (this *TOAConfig) SockFile() string {
	if len(this.SockPath) == 0 {
		return "/tmp/edge-toa.sock"
	}
	return this.SockPath
}

// 获取随机端口
func (this *TOAConfig) RandLocalPort() uint16 {
	return uint16(rands.Int(this.minLocalPort, this.maxLocalPort))
}