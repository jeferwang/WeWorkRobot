package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerConfig *ServerConfig `json:"server"`
	WxWorkConfig *WxWorkConfig `json:"wxwork"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type WxWorkConfig struct {
	Token          string `json:"Token"`
	EncodingAESKey string `json:"EncodingAESKey"`
}

var config *Config

func Setup(confPath string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fullPath := path.Join(cwd, confPath)
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}
	config = new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}
	return nil
}

func GetServerConfig() *ServerConfig {
	return config.ServerConfig
}

func GetWxWorkConfig() *WxWorkConfig {
	return config.WxWorkConfig
}

func (c *ServerConfig) GetListenAddr() string {
	var buf bytes.Buffer
	buf.WriteString(c.Host)
	buf.WriteString(":")
	buf.WriteString(c.Port)
	return buf.String()
}

func (c *ServerConfig) GetMode() string {
	m := strings.ToLower(c.Mode)
	if m == "release" {
		return gin.ReleaseMode
	}
	if m == "test" {
		return gin.TestMode
	}
	return gin.DebugMode
}
