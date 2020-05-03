package elastic

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/fpapadopou/music-index/internal/app/index"
	"io"
	"time"
)

type document struct {
	ID        string `json:"_id"`
	Timestamp int64  `json:"timestamp"`
	Term      string `json:"term"`
	Type      string `json:"type"`
	Name      string `json:"name"`
}

func (d document) generateID() (string, error) {

	h := md5.New()
	data := fmt.Sprintf("term:%s_type:%s_name:%s", d.Term, d.Type, d.Name)
	_, err := io.WriteString(h, data)
	if err != nil {
		return "", err
	}

	return string(hex.EncodeToString(h.Sum(nil))), nil
}

func createDoc(s index.Suggestion) (*document, error) {
	doc := &document{
		Timestamp: time.Now().Unix(),
		Term:      s.Term,
		Type:      s.Type,
		Name:      s.Name,
	}

	ID, err := doc.generateID()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot create ES doc ID: %v", err))
	}

	doc.ID = ID
	return doc, nil
}
