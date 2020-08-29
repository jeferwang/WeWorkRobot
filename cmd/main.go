package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jeferwang/WeWorkRobot/src/conf"
	"github.com/jeferwang/WeWorkRobot/src/router"
)

func main() {
	Setup()
	gin.SetMode(conf.GetServerConfig().GetMode())
	r := gin.Default()
	router.Setup(r)
	_ = r.Run(conf.GetServerConfig().GetListenAddr())
}

func Setup() {
	var configPath string
	flag.StringVar(&configPath, "c", "", "配置文件路径")
	flag.Parse()
	err := conf.Setup(configPath)
	if err != nil {
		fmt.Printf("配置文件加载失败：%s\n", err.Error())
		return
	}
}
