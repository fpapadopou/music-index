package elastic

import (
	"context"
	"github.com/fpapadopou/music-index/internal/app/index"
)

// Client is a concrete implementation of the SuggestionsService interface using ElasticSearch.
type Client struct {
}

// TrackResults implements result tracking in ElasticSearch.
func (c *Client) TrackResults(ctx context.Context, q index.Query, rr []*index.Result) error {

	return nil
}

// Suggest implements suggestions retrieval from ElasticSearch.
func (c *Client) Suggest(ctx context.Context, q index.Query) ([]*index.Suggestion, error) {

	return []*index.Suggestion{}, nil
}
