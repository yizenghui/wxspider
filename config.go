package wxspider

import "github.com/BurntSushi/toml"

// Config 配置
type Config struct {
	BaiDuAiConf BaiDuAiConf
	PostConfig  PostConfig
	WebServe    WebServe
}

//BaiDuAiConf 百度ai的密令
type BaiDuAiConf struct {
	APIKey    string `toml:"api_key"`
	SecretKey string `toml:"secret_key"`
}

//WebServe 本地web服务器设置
type WebServe struct {
	Port int64 `toml:"port"`
}

//PostConfig 配置
type PostConfig struct {
	ServeURL           string `toml:"serve_url"`
	AuthorizationToken string `toml:"authorization_token"`
}

var confFile = "conf.toml"
var config Config

func init() {
	GetConf()
}

//GetConf 获取config
func GetConf() Config {
	if config.PostConfig.ServeURL == "" {
		if _, err := toml.DecodeFile(confFile, &config); err != nil {
			panic(err)
		}
	}
	return config
}
