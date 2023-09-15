package main

import (
	"context"
	"dubbo-go-client/api"
	"flag"
	"os"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var grpcGenerator1Impl = new(api.Generator1ClientImpl)

var grpcXTestImpl = new(api.XtestClientImpl) // 确保在使用之前进行初始化

func init() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "./conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)

	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := config.Load(); err != nil {
		panic(err)
	}
}

func main() {
	config.SetConsumerService(grpcGenerator1Impl)
	config.SetConsumerService(grpcXTestImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}

	logger.Info("start to test dubbo")
	req := &api.GenReq{
		AppId:  "laurence",
		AppAge: 30,
	}
	reply, err := grpcGenerator1Impl.GetID(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("get result~~: %v\n", reply)

	req2 := &api.XtestReq{
		Name: "xiongxiong~",
	}
	reply2, err := grpcXTestImpl.GetUser(context.Background(), req2)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("get result2~~: %v\n", reply2)
}
