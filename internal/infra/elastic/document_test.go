package elastic

import (
	"github.com/fpapadopou/music-index/internal/app/index"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_generateID(t *testing.T) {

	doc := document{
		Term: "may",
		Type: "artist",
		Name: "John Mayer",
	}

	hash, err := doc.generateID()
	assert.Equal(t, "39e55288a0f83c834f86bd6eb02db205", hash)
	assert.NoError(t, err)
}

func Test_createDoc(t *testing.T) {

	s := index.Suggestion{
		Term: "search me",
		Type: "song",
		Name: "Neon",
	}

	now := time.Now().Unix()
	doc, err := createDoc(s)

	assert.NoError(t, err)
	assert.Equal(t, "ae1a2f4cc9dcd1d2936cf9fc234efb81", doc.ID)
	assert.Equal(t, "search me", doc.Term)
	assert.Equal(t, "song", doc.Type)
	assert.Equal(t, "Neon", doc.Name)
	assert.InDelta(t, now, doc.Timestamp, 1)
}
