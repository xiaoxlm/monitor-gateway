package model

import (
	"github.com/xiaoxlm/monitor-gateway/config"
	"gorm.io/datatypes"
)

func init() {
	config.Config.Mysql.AppendModel(&MetricsMapping{})
}

type MetricsMapping struct {
	ID
	MetricUniqueID string            `json:"metricUniqueID" gorm:"unique"` // 告警唯一标识
	Labels         datatypes.JSONMap `json:"labels"`                       // 指标标签
	Expression     string            `json:"expression"`                   // 表达式
	Desc           string            `json:"description"`                  // 描述
	TimeAt
}

func (MetricsMapping) TableName() string {
	return "metrics_mapping"
}
