//go:generate mockgen -package=elastic -self_package=github.com/fpapadopou/music-index/internal/infra/elastic/elastic -destination=../es_transport_mock_test.go github.com/elastic/go-elasticsearch/v7/esapi Transport

package transport

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

// New returns a new ES client implementing the esapi.Transport interface.
func New() (esapi.Transport, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	c, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	log.Println(c.Info())
	return c, nil
}
