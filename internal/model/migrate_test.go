package model

import (
	"github.com/xiaoxlm/monitor-gateway/config"
	"testing"
)

func TestMigrate(t *testing.T) {
	err := config.Config.Mysql.MigrateTable(&MetricsMapping{})
	if err != nil {
		t.Fatal(err)
	}
}
