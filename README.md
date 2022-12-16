# airmart_pay项目说明
---

## 项目结构
- airmart_pay 项目服务
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
go env -w GOPRIVATE=codeup.aliyun.com
go env -w GONOPROXY=codeup.aliyun.com
git config --global url."ssh://git@codeup.aliyun.com".insteadOf "https://codeup.aliyun.com"
```

## 快速开始
```shell
git clone git@codeup.aliyun.com:62da4ef41a358b4399af6f42/airmartServices/airmart_pay.git
cd airmart_pay
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
Nacos  配置中心（viper）
JWT    鉴权
Logger zap(考虑集成logrus)
Kafka  消息队列
OSS    对象存储（bucket）
ElasticSearch 大数据处理（基本实现所有Document APIs,还未集成到airmart-core）
```