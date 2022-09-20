package service

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearchConfig struct{}

var (
	ecurl = "http://elasticsearch:9200" //viper.Get("SERVICE.ELASTICSEARCH_URL").(string)
)

var cfg = elasticsearch.Config{
	Addresses: []string{
		ecurl,
	},
}
var es, _ = elasticsearch.NewClient(cfg)

func NewElasticSearch() *ElasticSearchConfig {
	return &ElasticSearchConfig{}
}

func (*ElasticSearchConfig) GetInfo() error {
	fmt.Println(es.Info())
	return nil
}
