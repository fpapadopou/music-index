// +build integration

package elastic

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/fpapadopou/music-index/internal/app/index"
	"github.com/fpapadopou/music-index/internal/infra/elastic/transport"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var es esapi.Transport

func TestMain(m *testing.M) {
	// Create Resource pool.
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	// Pulls an image, creates a container and runs it.
	opts := &dockertest.RunOptions{
		Repository: "docker.elastic.co/elasticsearch/elasticsearch",
		Tag:        "7.6.2",
		Env: []string{
			"node.name=es_int_test",
			"cluster.name=es-int-test-cluster",
			"cluster.initial_master_nodes=es_int_test",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("9200/tcp"): []docker.PortBinding{docker.PortBinding{HostIP: "0.0.0.0", HostPort: "9200"}},
		},
	}
	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		log.Fatalf("could not start Elastic resource: %s", err)
	}

	// Connect to ES, retry with exponential delay.
	if err = pool.Retry(func() error {
		var err error
		es, err = transport.New()

		return err
	}); err != nil {
		log.Fatalf("could not connect to ES container: %s", err)
	}

	code := m.Run()

	// Fail if the resource cannot be purged - will have to be removed manually.
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("could not purge ES container: %s", err)
	}

	os.Exit(code)
}

func TestClient_TrackResults_Integration(t *testing.T) {
	/**
	Happy path test. Index `suggestion` is auto-created with the current setup when
	the first document is pushed to ES.
	*/
	client := New(es)

	q := index.Query{
		Type: "artist",
		Term: "may",
	}

	rr := []*index.Result{
		{Type: "artist", Name: "John Mayer"},
		{Type: "artist", Name: "Mayhem"},
	}

	err := client.TrackResults(context.Background(), q, rr)
	assert.NoError(t, err)
}
