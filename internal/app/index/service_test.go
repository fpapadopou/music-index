package index

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Query_Search_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	service, err := New(searchService, suggestionsService)
	assert.NoError(t, err)

	q := Query{}
	searchService.
		EXPECT().
		Search(context.TODO(), q).
		Return(nil, errors.New("error in search"))

	res, err := service.Query(context.TODO(), q)

	assert.Empty(t, res)
	assert.Error(t, err, "error in search")
}

func TestService_Query_Search_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	service, err := New(searchService, suggestionsService)
	assert.NoError(t, err)

	q := Query{}
	expected := []*Result{
		{Name: "Some title!", Type: "Song"},
	}

	searchService.
		EXPECT().
		Search(context.TODO(), q).
		Return(expected, nil)

	// TODO: Fix test for goroutine.
	//suggestionsService.
	//	EXPECT().
	//	TrackResults(context.TODO(), q, expected).
	//	Return(nil)

	res, err := service.Query(context.TODO(), q)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, res)
}

func TestService_Suggest_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	service, err := New(searchService, suggestionsService)
	assert.NoError(t, err)

	q := Query{}
	expected := []*Suggestion{
		{Name: "Some guy!", Type: "Artist"},
	}

	suggestionsService.
		EXPECT().
		Suggest(context.TODO(), q).
		Return(expected, nil)

	res, err := service.Suggest(context.TODO(), q)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, res)
}

func TestNew(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	searchService := NewMockSearchClient(controller)
	suggestionsService := NewMockSuggestionsClient(controller)

	_, err := New(searchService, suggestionsService)
	assert.NoError(t, err)
}
