package model

import (
	"github.com/xiaoxlm/monitor-gateway/config"
	"github.com/xiaoxlm/monitor-gateway/internal/enum"
	"gorm.io/datatypes"
)

func init() {
	config.Config.Mysql.AppendModel(&MetricsMapping{})
}

type MetricsMapping struct {
	ID
	MetricUniqueID enum.MetricUniqueID         `json:"metricUniqueID" gorm:"unique"` // 告警唯一标识
	Labels         datatypes.JSONMap           `json:"labels"`                       // 指标标签(key:标签名；value:标签描述)
	Expression     string                      `json:"-"`                            // 表达式
	Desc           string                      `json:"description"`                  // 描述
	Category       enum.MetrcisMappingCategory `json:"category"`                     // 类别
	TimeAt
}

func (MetricsMapping) TableName() string {
	return "metrics_mapping"
}
