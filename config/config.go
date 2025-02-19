package config

import (
	"fmt"
	tool "github.com/lie-flat-planet/service-init-tool"
	"github.com/lie-flat-planet/service-init-tool/component/mysql"
)

func init() {
	aa := Config
	fmt.Print(aa)

	if err := tool.Init("./", &Config); err != nil {
		panic(err)
	}
}

var Config = struct {
	Server *tool.Server
	Mysql  *mysql.Mysql
}{
	Server: &tool.Server{
		Name:     "monitor-gateway",
		Code:     333 * 1e3,
		HttpPort: 8081,
		RunMode:  "debug",
	},
	Mysql: &mysql.Mysql{
		Config: mysql.Config{
			Host:        "10.10.1.84:3306",
			User:        "root",
			Password:    "uWXf87plmQGz8zMM",
			DbName:      "n9e_v6",
			MaxIdleConn: 1,
			MaxOpenConn: 2,
		},
	},
}
