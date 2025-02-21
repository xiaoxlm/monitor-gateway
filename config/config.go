package config

import (
	tool "github.com/lie-flat-planet/service-init-tool"
	"github.com/lie-flat-planet/service-init-tool/component/mysql"
	"github.com/lie-flat-planet/service-init-tool/component/prometheus"
)

func init() {
	if err := tool.Init("./", &Config); err != nil {
		panic(err)
	}
}

var Config = struct {
	ClientID     string `env:""`
	ClientSecret string `env:""`
	Server       *tool.Server
	Mysql        *mysql.Mysql
	Prom         *prometheus.Prom
}{
	ClientID:     "xx",
	ClientSecret: "222",
	Server: &tool.Server{
		Name:     "monitor-gateway",
		Code:     333 * 1e3,
		HttpPort: 8081,
		RunMode:  "debug",
	},
	Mysql: &mysql.Mysql{
		Config: mysql.Config{
			Host:        "127.0.0.1:3306",
			User:        "root",
			Password:    "11111",
			DbName:      "2222",
			MaxIdleConn: 1,
			MaxOpenConn: 2,
		},
	},
	Prom: &prometheus.Prom{
		Addr: "1.2.3.4",
	},
}
