package _interface

import (
	"context"

	"github.com/prometheus/common/model"
)

type QueryFormItem struct {
	Start int64  `json:"start"` // 开始时间
	End   int64  `json:"end"`   // 结束时间
	Step  uint   `json:"step"`  // 步长
	Query string `json:"query"` // 查询语句
}

type TimeSeriesDB interface {
	QueryRange(ctx context.Context, queries QueryFormItem) (model.Value, error)
}
