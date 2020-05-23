//go:generate mockgen -package=elastic -self_package=github.com/fpapadopou/music-index/internal/infra/elastic/elastic -destination=../es_transport_mock_test.go github.com/elastic/go-elasticsearch/v7/esapi Transport

package transport

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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

	log.Printf("init ES client (version %v)", elasticsearch.Version)
	res, err := c.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting ES transport info: %+v, %+v", res, err)
	}
	log.Printf("ES info: %+v", res)

	defer res.Body.Close()
	return c, nil
}
