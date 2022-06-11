package models

import (
	"context"
	"encoding/json"
	"testing"
	"time"
	//"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testIndex = "test-index"
)

func initTestESClient() (*elasticsearch.Client, error) {
	esclient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	return esclient, err
}

func TestCreateIndex(t *testing.T) {
	esclient, err := initTestESClient()
	assert.NoError(t, err)

	elastic := &Elastic{
		client:  esclient,
		index:   testIndex,
		timeout: 3 * time.Second,
	}

	err = elastic.CreateIndex(testIndex)
	require.NoError(t, err)
}

func TestCreateResource(t *testing.T) {
	esclient, err := initTestESClient()
	assert.NoError(t, err)

	elastic := &Elastic{
		client:  esclient,
		index:   testIndex,
		timeout: 3 * time.Second,
	}
	newResource := Resource{
		ID:        "test1",
		Title:     "test1-title",
		Author:    "-",
		Content:   "some-content",
		CreatedAt: time.Now().Format(time.UnixDate),
		Tags:      []string{"test"},
	}
	err = elastic.CreateResource(newResource)

	req := esapi.GetRequest{
		Index:      testIndex,
		DocumentID: "test1",
	}

	res, err := req.Do(context.Background(), esclient)
	require.NoError(t, err)

	require.False(t, res.IsError())

	var response GetResponseModel
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	require.True(t, response.Found)

	require.Equal(t, newResource.ID, response.Source.ID)
	require.Equal(t, newResource.Author, response.Source.Author)
	require.Equal(t, newResource.Content, response.Source.Content)
	require.Equal(t, newResource.Title, response.Source.Title)
	require.Equal(t, newResource.Tags, response.Source.Tags)
	require.Equal(t, newResource.CreatedAt, response.Source.CreatedAt)
}

type GetResponseModel struct {
	Index       string   `json:"index"`
	ID          string   `json:"id"`
	Version     int      `json:"_version"`
	SeqNo       int      `json:"_seq_no"`
	PrimaryTerm int      `json:"_primary_term"`
	Found       bool     `json:"found"`
	Source      Resource `json:"_source"`
}

func TestUpdateResource(t *testing.T) {
	esclient, err := initTestESClient()
	assert.NoError(t, err)

	elastic := &Elastic{
		client:  esclient,
		index:   testIndex,
		timeout: 3 * time.Second,
	}

	const id = "test1"

	resourceToUpdate := ResourceData{
		Title:     "updated title",
		Author:    "-",
		Content:   "some new updated content",
		Tags:      []string{"test", "some-tag"},
	}

	err = elastic.UpdateResource(id, resourceToUpdate)
	require.NoError(t, err)

	req := esapi.GetRequest{
		Index:      testIndex,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), esclient)
	require.NoError(t, err)

	require.False(t, res.IsError())

	response := map[string]interface{}{
		"_source":make(map[string]interface{}),
	}


	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	

	require.True(t, response["found"].(bool))


}