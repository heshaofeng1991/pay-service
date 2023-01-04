# pay-service项目说明
---

## 项目结构
- pay-service 项目服务
  - config 服务配置文件位置
  - global 服务全局配置初始化
  - internal 服务内部使用的一些东西
  - model 服务持久化层model定义(数据库表结构定义)
  - pay 各个支付的集合
    - whole_pay 全民付
    - other_pay 其它支付
  - router 服务路由
  - service 服务持久化层具体实现
  - types 服务出入参定义
  - main.go 服务入口
  - Makefile shell脚本命令

## 项目依赖设置 go mod && git config
```shell
go env -w GO111MODULE=on
```

## 快速开始
```shell
git clone git@github.com:heshaofeng1991/pay-service.git
cd pay-service
go mod tidy -compat=1.9
go run main.go -confFile ./config/config-dev.yaml
```

## 其它组件
```text
Go1.19 Golang最新版本
Consul 服务注册，服务发现
Jaeger Trace链路追踪，监控
Mysql  数据库
Redis  Redis
Nacos  配置中心
JWT    鉴权
Logger zap
Kafka  消息队列
OSS    对象存储
ElasticSearch
```