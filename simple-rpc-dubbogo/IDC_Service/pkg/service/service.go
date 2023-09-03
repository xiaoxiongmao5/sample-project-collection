package service

import (
	"context"
	"dubbo-go-app/api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	uuid "github.com/satori/go.uuid"
)

type Generator1ServerImpl struct {
	api.UnimplementedGenerator1Server
}

func (s *Generator1ServerImpl) GetID(ctx context.Context, in *api.GenReq) (*api.GenResp, error) {
	logger.Infof("Dubbo-go GeneratorProvider AppId = %s\n", in.AppId)
	uuid, err := uuid.NewV4()
	if err != nil {
		logger.Infof("Dubbo-go GeneratorProvider get id err = %v\n", err)
		return nil, err
	}
	return &api.GenResp{Id: uuid.String(), Name: "xiongxiong", Age: in.AppAge}, nil
}

type XtestServerImpl struct {
	api.UnimplementedXtestServer
}

func (s *XtestServerImpl) GetUser(ctx context.Context, in *api.XtestReq) (*api.XtestResp, error) {
	logger.Infof("Dubbo-go XtestProvider Name = %s\n", in.Name)
	return &api.XtestResp{
		Name: in.Name,
		Age:  18,
		Sex:  "woman",
	}, nil
}

func init() {
	config.SetProviderService(&Generator1ServerImpl{})
	config.SetProviderService(&XtestServerImpl{})
}
