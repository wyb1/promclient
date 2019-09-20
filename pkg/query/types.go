package query

import (
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// Querier can use prometheus API to send queries directly to Prometheus
type Querier struct {
	client api.Client
	api    v1.API
}
