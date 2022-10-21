package repositories

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

type BigQuery struct {
	client *bigquery.Client
}

// New returns a new BigQuery instance.
func NewBigQuery(ctx context.Context, serviceAccountRaw string) (*BigQuery, error) {
	var bq BigQuery

	client, err := bigquery.NewClient(ctx, "DetectProjectID", option.WithCredentialsJSON([]byte(serviceAccountRaw)))
	if err != nil {
		return nil, err
	}

	bq.client = client
	return &bq, nil
}
