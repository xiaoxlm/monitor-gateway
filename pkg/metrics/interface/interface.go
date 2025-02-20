package _interface

import (
	"context"
)

type QueryFormItem struct {
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Step  int64  `json:"step"`
	Query string `json:"query"`
}

type Metrics interface {
	BatchQueryRange(ctx context.Context, queries []QueryFormItem) error
	Marshal() ([]byte, error)
	Unmarshal(data []byte, v any) error
}
