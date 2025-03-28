package prometheus

import (
	"context"
	_interface "github.com/xiaoxlm/monitor-gateway/pkg/metrics/interface"
	"time"

	"github.com/prometheus/client_golang/api"
	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type Prometheus struct {
	client api.Client
}

func NewPrometheus(cli api.Client) (*Prometheus, error) {
	return &Prometheus{client: cli}, nil
}

func (p *Prometheus) QueryRange(ctx context.Context, item _interface.QueryFormItem) (model.Value, error) {
	r := prometheus_v1.Range{
		Start: time.Unix(item.Start, 0),
		End:   time.Unix(item.End, 0),
		Step:  time.Duration(item.Step) * time.Second,
	}

	resp, _, err := prometheus_v1.NewAPI(p.client).QueryRange(ctx, item.Query, r)
	return resp, err
}
