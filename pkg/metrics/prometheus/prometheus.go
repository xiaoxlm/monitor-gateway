package prometheus

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/api"
	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type QueryFormItem struct {
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Step  int64  `json:"step"`
	Query string `json:"query"`
}

type Prometheus struct {
	client api.Client
}

func NewPrometheus(addr string) (*Prometheus, error) {
	conf := api.Config{
		Address: addr,
	}

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &Prometheus{client: client}, nil
}

func (p *Prometheus) BatchQueryRange(ctx context.Context, queries []QueryFormItem) ([]model.Value, error) {
	var lists []model.Value

	for _, item := range queries {
		r := prometheus_v1.Range{
			Start: time.Unix(item.Start, 0),
			End:   time.Unix(item.End, 0),
			Step:  time.Duration(item.Step) * time.Second,
		}

		resp, _, err := prometheus_v1.NewAPI(p.client).QueryRange(ctx, item.Query, r)
		if err != nil {
			return nil, err
		}

		lists = append(lists, resp)
	}

	return lists, nil
}
