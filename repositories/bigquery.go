package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/Reskill-2022/hoarder/config"
	"github.com/Reskill-2022/hoarder/env"
	"github.com/Reskill-2022/hoarder/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type BigQuery struct {
	conf   config.Config
	client *bigquery.Client
}

// New returns a new BigQuery instance.
func NewBigQuery(ctx context.Context, serviceAccountRaw string, conf config.Config) (*BigQuery, error) {
	var bq BigQuery

	client, err := bigquery.NewClient(ctx, bigquery.DetectProjectID, option.WithCredentialsJSON([]byte(serviceAccountRaw)))
	if err != nil {
		return nil, err
	}

	bq.conf = conf
	bq.client = client
	return &bq, nil
}

func (bq *BigQuery) CreateSlackMessage(ctx context.Context, message models.SlackMessage) error {
	inserter := bq.client.Dataset(bq.conf.GetString(env.BigQuerySlackDatasetID)).
		Table(bq.conf.GetString(env.BigQuerySlackTableID)).
		Inserter()
	return inserter.Put(ctx, message)
}

func (bq *BigQuery) CreateZendeskTicket(ctx context.Context, ticket models.ZendeskTicket) error {
	inserter := bq.client.Dataset(bq.conf.GetString(env.BigQueryZendeskDatasetID)).
		Table(bq.conf.GetString(env.BigQueryZendeskTableID)).
		Inserter()
	return inserter.Put(ctx, ticket)
}

func (bq *BigQuery) CreateCalendlyEvent(ctx context.Context, event models.CalendlyEvent) error {
	inserter := bq.client.Dataset(bq.conf.GetString(env.BigQueryCalendlyDatasetID)).
		Table(bq.conf.GetString(env.BigQueryCalendlyTableID)).
		Inserter()
	return inserter.Put(ctx, event)
}

func (bq *BigQuery) CreateMoodleLogLine(ctx context.Context, line models.MoodleLogLine) error {
	inserter := bq.client.Dataset(bq.conf.GetString(env.BigQueryMoodleDatasetID)).
		Table(bq.conf.GetString(env.BigQueryMoodleLogsTableID)).
		Inserter()
	return inserter.Put(ctx, line)
}

// GetLastMoodleLogLine returns the last Moodle log line sorted by timecreated.
func (bq *BigQuery) GetLastMoodleLogLine(ctx context.Context) (*models.MoodleLogLine, error) {
	var line models.MoodleLogLine

	canonicalTableID := fmt.Sprintf("`%s.%s.%s`",
		bq.conf.GetString(env.BigQueryProjectID),
		bq.conf.GetString(env.BigQueryMoodleDatasetID),
		bq.conf.GetString(env.BigQueryMoodleLogsTableID),
	)
	query := bq.client.Query("SELECT * FROM " + canonicalTableID + " ORDER BY time_created DESC LIMIT 1")

	it, err := query.Read(ctx)
	if err != nil {
		return nil, err
	}

	for {
		err := it.Next(&line)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return &line, nil
}
