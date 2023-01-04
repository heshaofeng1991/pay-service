package main

import (
	"core/common"
	"core/initialize"
	"pay-service/global"
	"pay-service/router"

	"go.uber.org/zap"
)

func main() {
	// 服务配置初始化
	err := initialize.Config(true, global.Config)
	if err != nil {
		zap.S().Fatalf("加载配置失败 err:%v", err.Error())
	}

	zap.S().Infof("加载配置成功")
	// 默认依赖组件初始化
	global.Srv, err = initialize.DefaultServices()
	if err != nil {
		zap.S().Fatalf("依赖服务初始化失败 err:%v", err.Error())
	}

	zap.S().Infof("依赖服务初始化成功")

	// http服务启动
	common.Run(router.Init(), global.Srv)
}
