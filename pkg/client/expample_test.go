package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestBatchQuery(t *testing.T) {
	client, err := NewClientWithResponses("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	// Example: Query metrics
	ctx := context.Background()

	// Create request body
	body := BatchQueryJSONRequestBody{
		Queries: []GithubComXiaoxlmMonitorGatewayApiRequestMetricsQueryInfo{
			{
				MetricUniqueID: "cpu_avg_util",
				Start:          1742178824,
				End:            1742178824,
				Step:           10,
				LabelValue: map[string]string{
					"IBN":     "算网A",
					"host_ip": "10.10.1.85",
				},
			},
		},
	}

	// Add authorization header
	headers := &BatchQueryParams{
		Authorization: "Basic dG9uZzp0b25nVGVjaEAyMDI1",
	}

	// Make the request
	resp, err := client.BatchQueryWithResponse(ctx, headers, body)
	if err != nil {
		t.Fatal(err)
	}

	// success
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		LogJSON(*resp.JSON200.Data)
		return
	}

	if resp.JSON500 != nil {
		// process 500 data
		LogJSON(*resp.JSON500)
		return
	}

	if resp.JSON400 != nil {
		// process 400 data
		LogJSON(*resp.JSON400)
		return
	}
}

func LogJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
