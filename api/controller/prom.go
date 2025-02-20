package controller

//
//type PromController struct {
//	prom *prometheus.Prometheus
//}
//
//func NewPromController(promAddr string) (*PromController, error) {
//	prom, err := prometheus.NewPrometheus(promAddr)
//	if err != nil {
//		return nil, err
//	}
//
//	return &PromController{prom: prom}, nil
//}
//
//func (p *PromController) BatchQueryRange(ctx context.Context, queries []prometheus.QueryFormItem) ([]model.Value, error) {
//	return p.prom.BatchQueryRange(ctx, queries)
//}
