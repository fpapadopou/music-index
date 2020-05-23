package elastic

import (
	"context"
	"errors"
	"github.com/fpapadopou/music-index/internal/app/index"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_TrackResults_ClientErrorNoResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	tr := NewMockTransport(ctrl)

	tr.EXPECT().
		Perform(gomock.Any()).
		Return(nil, errors.New("es transport error (body might be missing)")).
		MaxTimes(1)

	client := New(tr)

	q := index.Query{
		Type: "artist",
		Term: "may",
	}

	rr := []*index.Result{
		{Type: "artist", Name: "John Mayer"},
	}

	err := client.TrackResults(context.Background(), q, rr)
	assert.Error(t, err)
}

func TestClient_TrackResults_FailureResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	tr := NewMockTransport(ctrl)

	resp := &http.Response{
		StatusCode: 500,
		Body:       ioutil.NopCloser(strings.NewReader("internal server error")),
	}

	tr.EXPECT().
		Perform(gomock.Any()).
		Return(resp, nil).
		MaxTimes(2)

	client := New(tr)

	q := index.Query{
		Type: "song",
		Term: "bla",
	}

	rr := []*index.Result{
		{Type: "song", Name: "Blak and Blue"},
		{Type: "artist", Name: "Gary Clark Jr."},
	}

	err := client.TrackResults(context.Background(), q, rr)
	assert.Error(t, err)
}

func TestClient_TrackResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	tr := NewMockTransport(ctrl)

	resp := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader("suggestion indexed in ES")),
	}

	tr.EXPECT().
		Perform(gomock.Any()).
		Return(resp, nil).
		MaxTimes(2)

	client := New(tr)

	q := index.Query{
		Type: "song",
		Term: "bla",
	}

	rr := []*index.Result{
		{Type: "song", Name: "Blak and Blue"},
		{Type: "artist", Name: "Gary Clark Jr."},
	}

	err := client.TrackResults(context.Background(), q, rr)
	assert.NoError(t, err)
}
