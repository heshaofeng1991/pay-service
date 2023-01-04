package global

import (
	conf "core/config"

	"core/initialize"
)

type config struct {
	Mysql   conf.DB      `json:"mysql"`
	Redis   conf.Redis   `json:"redis"`
	Service conf.Service `json:"service"`
	Jaeger  conf.Jaeger  `json:"jaeger"`
	Consul  conf.Consul  `json:"consul"`
	Logs    conf.Logs    `json:"logs"`
	Master  MasterJWT    `json:"master"`
}

type MasterJWT struct {
	Jwt  conf.Jwt `json:"jwt"`
	Host string   `json:"host"`
	Auth string   `json:"auth"`
}

var (
	Config *config
	Srv    *initialize.Services
)

func init() {
	Config = new(config)
}
