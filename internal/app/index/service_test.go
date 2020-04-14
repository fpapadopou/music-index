package index

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Query_Search(t *testing.T) {
	var searchTests = []struct {
		n             string
		c             context.Context
		q             Query
		expectResults []*Result
		expectError   error
	}{
		{"success", context.TODO(), Query{}, []*Result{{Name: "Title", Type: "song"}}, nil},
		{"failure", context.TODO(), Query{}, nil, errors.New("failed")},
	}

	for _, tt := range searchTests {
		t.Run(tt.n, func(t *testing.T) {
			s, ctrl, ss, _ := createService(t)
			defer ctrl.Finish()

			ss.EXPECT().Search(tt.c, tt.q).Return(tt.expectResults, tt.expectError)

			res, err := s.Query(tt.c, tt.q)

			assert.Equal(t, res, tt.expectResults, fmt.Sprintf("got %+v, want %+v", tt.expectResults, res))
			assert.Equal(t, err, tt.expectError, fmt.Sprintf("got %q, want %q", tt.expectError, err))
		})
	}
}

func TestService_Track(t *testing.T) {

	var trackTests = []struct {
		n      string
		c      context.Context
		q      Query
		rr     []*Result
		expect error
	}{
		{"success", context.TODO(), Query{}, []*Result{{Name: "Title", Type: "song"}}, nil},
		{"failure", context.TODO(), Query{}, []*Result{{Name: "Title", Type: "song"}}, errors.New("failed")},
	}

	for _, tt := range trackTests {
		t.Run(tt.n, func(t *testing.T) {
			s, ctrl, _, ss := createService(t)
			defer ctrl.Finish()

			ss.EXPECT().TrackResults(tt.c, tt.q, tt.rr).Return(tt.expect)

			err := s.Track(tt.c, tt.q, tt.rr)

			assert.Equal(t, err, tt.expect, fmt.Sprintf("got %q, want %q", tt.expect, err))
		})
	}
}

func TestService_Suggest(t *testing.T) {
	var suggestTests = []struct {
		n                string
		c                context.Context
		q                Query
		expectSuggestion []*Suggestion
		expectError      error
	}{
		{"success", context.TODO(), Query{}, []*Suggestion{{Name: "Hugh Laurie", Type: "artist"}}, nil},
		{"failure", context.TODO(), Query{}, nil, errors.New("failed")},
	}

	for _, tt := range suggestTests {
		t.Run(tt.n, func(t *testing.T) {
			s, ctrl, _, ss := createService(t)
			defer ctrl.Finish()

			ss.EXPECT().Suggest(tt.c, tt.q).Return(tt.expectSuggestion, tt.expectError)

			res, err := s.Suggest(tt.c, tt.q)

			assert.Equal(t, res, tt.expectSuggestion, fmt.Sprintf("got %+v, want %+v", tt.expectSuggestion, res))
			assert.Equal(t, err, tt.expectError, fmt.Sprintf("got %q, want %q", tt.expectError, err))
		})
	}
}

func TestNew(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	_, err := New(searchService, suggestionsService)
	assert.NoError(t, err)
}

func createService(t *testing.T) (*Service, *gomock.Controller, *MockSearchClient, *MockSuggestionsClient) {
	controller := gomock.NewController(t)
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	service, _ := New(searchService, suggestionsService)

	return service, controller, searchService, suggestionsService
}
