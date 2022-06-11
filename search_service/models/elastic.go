package models

import (
	"fmt"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Elastic struct {
	client  *elasticsearch.Client
	index   string
	timeout time.Duration
}

func NewESClient(addrs []string) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses:    addrs,
		RetryOnError: retryOnError,
	})
}

func NewElastic(addr []string, index string, timeout int64) *Elastic {
	esclient, err := NewESClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return &Elastic{
		client:  esclient,
		index:   index,
		timeout: time.Duration(timeout) * time.Second,
	}
}

func retryOnError(req *http.Request, err error) bool {

	return false
}

func (e *Elastic) CreateIndex(index string) error {
	res, err := e.client.Indices.Exists([]string{index})
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		return nil
	}

	res, err = e.client.Indices.Create(index)
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}

func (e *Elastic) CreateResource(resource Resource) error {
	data, err := json.Marshal(resource)
	if err != nil {
		return err
	}

	req := esapi.CreateRequest{
		Index:      e.index,
		DocumentID: resource.ID,
		Body:       bytes.NewReader(data),
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 409 {
		return errors.New("multiple users are trying to update the same version of the document at the same time")
	}

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}

func (e *Elastic) UpdateResource(id string, resource ResourceData) error {
	data, err := json.Marshal(resource)
	if err != nil {

	}

	req := esapi.UpdateRequest{
		Index:      e.index,
		DocumentID: id,
		Body:        bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, string(data)))),
	}
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("document not found")
	}

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}

func (e *Elastic) DeleteResource(id string) error {
	req := esapi.DeleteRequest{
		Index:      e.index,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		return err
	}

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}
