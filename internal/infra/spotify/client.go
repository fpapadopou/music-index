package spotify

import (
	"context"
	"github.com/fpapadopou/music-index/internal/app/index"
)

// Client is a concrete implementation of the SearchClient interface using Spotify's REST API.
type Client struct {
}

// Search queries Spotify's API and retrieves results (tracks, albums, artists).
func (c *Client) Search(ctx context.Context, q index.Query) ([]*index.Result, error) {

	return []*index.Result{}, nil
}
