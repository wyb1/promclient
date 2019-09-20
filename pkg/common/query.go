package common

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

// InitClient creates a new client for the querier
func InitClient(address string) (*Querier, error) {
	client, err := api.NewClient(api.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	api := v1.NewAPI(client)

	return &Querier{
		Client: client,
		Api:    api,
	}, err
}

// PrintQuery sends a string query to Prometheus
func (q *Querier) PrintQuery(ctx context.Context, query string) error {
	if q.Client == nil || q.Api == nil {
		return errors.New("Client not initialized")
	}

	result, warnings, err := q.Api.Query(ctx, query, time.Now())
	if err != nil {
		return err
	}

	fmt.Printf("Query: %s\n", query)
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	fmt.Printf("Result:\n%v\n", result)
	fmt.Println("-----------------------------------------")
	return nil
}

// GetQuery returns the model of a given query
func (q *Querier) GetQuery(ctx context.Context, query string) (*model.Value, api.Warnings, error) {
	return nil, nil, nil
}
