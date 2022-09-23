package service

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/spf13/viper"
)

type ElasticSearchClient struct {
	host   string
	client *elastic.Client
	err    error
}

func NewElasticSearch() *ElasticSearchClient {
	host := viper.Get("SERVICE.ELASTICSEARCH_URL").(string)
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host), elastic.SetBasicAuth("elastic", "changeme"))
	if err != nil {
		panic(err)
	}
	return &ElasticSearchClient{
		host:   host,
		client: client,
		err:    err,
	}
}

func (c *ElasticSearchClient) Ping() (*elastic.PingResult, error) {
	ctx := context.Background()
	info, _, err := c.client.Ping(c.host).Do(ctx)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *ElasticSearchClient) HasIndex(idx string, ctx context.Context) bool {
	e, err := c.client.IndexExists(idx).Do(ctx)
	if err != nil {
		return false
	}
	if !e {
		return false
	}
	return true
}
