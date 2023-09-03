package main

import (
	"context"
	"dubbo-go-client/api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var grpcGenerator1Impl = new(api.Generator1ClientImpl)

var grpcXTestImpl = new(api.XtestClientImpl) // 确保在使用之前进行初始化

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
