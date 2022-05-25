package service

import (
	"bytes"
	"context"
	"errors"
	"encoding/json"
	"search_service/models"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Service struct{
	elastic *elasticsearch.Client
	timeout time.Duration
	index string
}

func NewService(client *elasticsearch.Client, index string, timeout time.Duration) *Service{
	return &Service{
		elastic: client,
		timeout: timeout,
		index: index,
	}
}

func (s *Service) GetResourceByID(id string) (*models.Resource, error){
	req := esapi.AsyncSearchGetRequest{DocumentID: id}
	res, err := req.Do(context.Background(), s.elastic)
	if err != nil{
		return nil, err
	}
	if res.IsError(){

	}

	if res.StatusCode == 404{

	}
	resource := &models.Resource{}
	err = json.NewDecoder(res.Body).Decode(resource)
	if err != nil{

	}
	return resource, nil
}

func (s *Service) SearchResourcesByFilter(filter map[string]interface{}) ([]*models.Resource, error){
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query":filter,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil{
		return nil, err
	}
	res, err := s.elastic.Search(
		s.elastic.Search.WithTimeout(s.timeout),
		s.elastic.Search.WithContext(context.Background()),
		s.elastic.Search.WithIndex(s.index),
		s.elastic.Search.WithBody(&buf),
		s.elastic.Search.WithPretty(),
	)
	if err != nil{
		return nil, err
	}

	if res.IsError(){
		return nil, errors.New(res.String())
	}

	if res.StatusCode == 404{

	}

	var result []*models.Resource
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil{

	}
	return result, nil
}