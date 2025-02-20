package _interface

import (
	"context"

	"github.com/prometheus/common/model"
)

type QueryFormItem struct {
	Start int64  `json:"start" binding:"required"` // 开始时间
	End   int64  `json:"end" binding:"required"`   // 结束时间
	Step  int64  `json:"step" binding:"required"`  // 步长
	Query string `json:"query" binding:"required"` // 查询语句
}

type TimeSeriesDB interface {
	BatchQueryRange(ctx context.Context, queries []QueryFormItem) ([]model.Value, error)
}
