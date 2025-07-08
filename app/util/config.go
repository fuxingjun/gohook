package util

import (
	"encoding/json"
	"os"
	"sync"
)

type FromConfig struct {
	Platform string `json:"platform"`
	Key      string `json:"key"`
}

type ToConfig struct {
	Platform string `json:"platform"`
	Webhook  string `json:"webhook"`
}

type ConfigItem struct {
	HookFrom   FromConfig `json:"from"`
	HookToList []ToConfig `json:"to"`
}

func newEmptyConfig() {
	emptyConfig := []ConfigItem{
		{
			HookFrom: FromConfig{
				Platform: "lark",
				Key:      "your-custom-key",
			},
			HookToList: []ToConfig{
				{
					Platform: "qywx",
					Webhook:  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx",
				},
			},
		},
	}
	// json.MarshalIndent 会生成带缩进的 JSON, 便于阅读。
	jsonData, err := json.MarshalIndent(emptyConfig, "", "  ")
	if err != nil {
		panic(err)
	}
	// 0600 所有者可读写, 其他用户无权限 0644 所有者可读写, 其他用户只读。
	err = os.WriteFile("config.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadConfig(configPath string) *[]ConfigItem {
	// 打开 JSON 文件
	if configPath == "" {
		configPath = "config.json" // 默认配置文件路径
	}
	if !PathExists(configPath) {
		// 新建示例文件
		newEmptyConfig()
		os.Exit(1)
	}
	file, err := os.Open(configPath)
	if err != nil {
		panic("无法打开配置文件: " + err.Error())
	}
	defer file.Close()

	// 读取文件内容
	data := make([]byte, 1024*1024) // 1MB 缓冲区
	n, err := file.Read(data)
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}
	// 解析 JSON 到结构体
	var config []ConfigItem
	if err := json.Unmarshal(data[:n], &config); err != nil {
		panic("解析 JSON 失败: " + err.Error())
	}
	return &config
}

var (
	appConfig  *[]ConfigItem
	configOnce sync.Once
)

func GetAppConfig(configPath string) *[]ConfigItem {
	if appConfig == nil {
		configOnce.Do(func() {
			appConfig = ReadConfig(configPath)
		})
	}
	return appConfig
}
