package entity

import (
	"github.com/lie-flat-planet/httputil"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	"sort"
)

type colorHandlerVO struct {
	panel       *model.Panel
	metricsData httputil.MetricsFromExpr
}

func (vo *colorHandlerVO) setMetricsDataColor() error {
	for i, data := range vo.metricsData {
		for j := range data.Values {
			if err := vo.setColorByMetricsValues(&data.Values[j]); err != nil {
				return err
			}
		}

		vo.metricsData[i] = data
	}

	return nil
}

func (vo *colorHandlerVO) setColorByMetricsValues(mValue *httputil.MetricsValues) error {
	var (
		color = "#3FC453" // 绿色
		steps = vo.sortThresholdsStepsByValue()
	)

	for _, step := range steps {
		if step.Value == nil {
			continue
		}

		v := *step.Value

		if mValue.Value >= v {
			color = step.Color
			break
		}
	}

	mValue.Color = color
	return nil
}

func (vo *colorHandlerVO) sortThresholdsStepsByValue() []model.Step {
	thresholdsSteps := vo.panel.Options.Thresholds.Steps

	sort.Slice(thresholdsSteps, func(i, j int) bool {
		if thresholdsSteps[i].Value == nil || thresholdsSteps[j].Value == nil {
			return false
		}
		return *thresholdsSteps[i].Value > *thresholdsSteps[j].Value
	})

	return thresholdsSteps
}
