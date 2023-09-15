package main

import (
	_ "dubbo-go-app/pkg/service"
	"flag"
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func init() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)

	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := config.Load(); err != nil {
		panic(err)
	}
}

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
