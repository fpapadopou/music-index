package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/fpapadopou/music-index/internal/app/index"
)

const esIndex = "suggestion"

// Client is a concrete implementation of the SuggestionsService interface using ElasticSearch.
type Client struct {
	es esapi.Transport
}

// TrackResults implements result tracking in ElasticSearch.
func (c *Client) TrackResults(ctx context.Context, q index.Query, rr []*index.Result) error {

	errChn := make(chan error)
	done := make(chan bool)
	var wg sync.WaitGroup

	for _, r := range rr {
		wg.Add(1)

		go func(ctx context.Context, q index.Query, r *index.Result, wg *sync.WaitGroup) {
			defer wg.Done()

			err := c.indexResult(ctx, q, r)
			if err != nil {
				// TODO: Add metrics for error types.
				errChn <- fmt.Errorf("indexing result for query failed: %w", err)
			}
		}(ctx, q, r, &wg)
	}

	var errMsgs []string
	go func() {
		for e := range errChn {
			errMsgs = append(errMsgs, e.Error())
		}
		done <- true
	}()
	wg.Wait()
	close(errChn)
	<-done

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "\n"))
	}

	return nil
}

// Suggest implements suggestions retrieval from ElasticSearch.
func (c *Client) Suggest(ctx context.Context, q index.Query) ([]*index.Suggestion, error) {

	return []*index.Suggestion{}, nil
}

// New creates a new ElasticSearch SuggestionsClient.
func New(transport esapi.Transport) *Client {

	return &Client{es: transport}
}

func (c Client) indexResult(ctx context.Context, q index.Query, r *index.Result) error {

	doc, err := createDoc(index.Suggestion{
		Term: q.Term,
		Name: r.Name,
		Type: r.Type,
	})
	if err != nil {
		return fmt.Errorf("cannot create suggestion doc: %v", err)
	}

	body, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("cannot encode suggestion doc: %v", err)
	}

	req := esapi.IndexRequest{
		Index:      esIndex,
		DocumentID: doc.ID,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), c.es)
	if err != nil {
		return fmt.Errorf("cannot index suggestion doc: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing suggestion doc: %+v", res.String())
	}

	return nil
}
