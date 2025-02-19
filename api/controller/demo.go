package controller

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	"github.com/xiaoxlm/monitor-gateway/internal/service"
)

func FactoryDemo() *Demo {
	return &Demo{
		service: service.NewDemo(
			service.WithMysql(config.Config.Mysql.GetDB()),
			//service.WithRedis(config.Config.Redis.GetClient()),
		),
	}
}

type Demo struct {
	service *service.Demo
}

func (demo *Demo) FetchFirst(ctx context.Context) (*model.User, error) {
	return demo.service.FetchFirst(ctx)
}
