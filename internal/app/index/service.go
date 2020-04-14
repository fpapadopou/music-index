//go:generate mockgen -source=service.go -destination=./service_mock_test.go -package=index -self_package=github.com/fpapadopou/music-index/internal/app/index

package index

import (
	"context"
	"log"
)

// Service implements the domain logic, searching and suggestions retrieval.
type Service struct {
	search  SearchClient
	suggest SuggestionsClient
}

// SearchClient interface provides search functionality.
type SearchClient interface {
	Search(ctx context.Context, q Query) ([]*Result, error)
}

// SuggestionsClient interface handles result tracking and suggestions retrieval.
type SuggestionsClient interface {
	TrackResults(ctx context.Context, q Query, rr []*Result) error
	Suggest(ctx context.Context, q Query) ([]*Suggestion, error)
}

// Query performs a request to the search service and tracks the results, before returning them.
func (s *Service) Query(ctx context.Context, q Query) ([]*Result, error) {
	res, err := s.search.Search(ctx, q)
	if err != nil {
		log.Printf("Search errored: %v", err)
		return nil, err
	}

	return res, nil
}

// Track handles result tracking in the SuggestionsService.
func (s *Service) Track(ctx context.Context, q Query, rr []*Result) error {
	err := s.suggest.TrackResults(ctx, q, rr)
	if err != nil {
		log.Printf("Result tracking errored: %v", err)
	}

	return err
}

// Suggest fetches search suggestions from the SuggestionsService.
func (s *Service) Suggest(ctx context.Context, q Query) ([]*Suggestion, error) {
	res, err := s.suggest.Suggest(ctx, q)
	if err != nil {
		log.Printf("Suggestion retrieval errored: %v", err)
		return nil, err
	}

	return res, nil
}

// New returns a new Service.
func New(searchClient SearchClient, suggestionClient SuggestionsClient) (*Service, error) {
	return &Service{search: searchClient, suggest: suggestionClient}, nil
}
