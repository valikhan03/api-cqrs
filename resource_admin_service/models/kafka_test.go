package models

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"resource_admin_service/tools"
)
var tests = []*Event{
	&Event{
		EventType: "CREATE",
		Resource: Resource{
			ID: "test-id",
			Author: "test-author",

		},	
	}, 
	&Event{
		EventType: "UPDATE",
		Resource: Resource{

		},
	}, 
	&Event{
		EventType: "DELETE",
		Resource: Resource{

		},
	},
}

func TestExecuteEvents(t *testing.T){
	eventChan := make(chan *Event)
	testProducer := initTestProducer()
	require.NotNil(t, testProducer)

	testKafka := NewKafka(testProducer, eventChan, tools.InitLogger("./logs/test.txt", "kafka: "))

	testConsumer, testConsumerGroupHandler := initTestConsumerGroup(t)
	require.NotNil(t, testConsumer)
	require.NotNil(t, testConsumerGroupHandler)

	

	wg := &sync.WaitGroup{}
	wg.Add(3)
	
	go func(){
		for _, tt := range tests{
			eventChan <- tt
		}
		defer wg.Done()
	}()

	go func(){
		testKafka.ExecuteEvents()		
		defer wg.Done()
	}()

	go func(){
		err := testConsumer.Consume(context.Background(), []string{"CREATE", "UPDATE", "DELETE"}, testConsumerGroupHandler)
		if !assert.NoError(t, err) {
			log.Println(err)
		}
		defer wg.Done()
	}()

	wg.Wait()
}

func initTestProducer() sarama.SyncProducer{
	addrs := []string{"localhost:9092"}
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Timeout = 5 * time.Second
	prod, err := sarama.NewSyncProducer(addrs, configs)
	if err != nil {
		log.Println(err)
		return nil
	}
	return prod
}

func initTestConsumerGroup(t *testing.T) (sarama.ConsumerGroup, *TestConsumerGroupHandler){
	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	
	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "", configs)
	if err != nil{
		return nil, nil
	}

	return consumerGroup, &TestConsumerGroupHandler{t: t}
}


type TestConsumerGroupHandler struct{
	t *testing.T
}

func (h *TestConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var event Event
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			return err
		}
		switch event.EventType{
		case "CREATE":
			if assert.Equal(h.t, tests[0].Resource, event.Resource){
				return nil
			}else{
				return fmt.Errorf("unexpected data: \n%v\n%v", tests[0].Resource, event.Resource)
			}
		case "UPDATE":
			if assert.Equal(h.t, tests[1].Resource, event.Resource){
				return nil
			}else{
				return fmt.Errorf("unexpected data: \n%v\n%v", tests[1].Resource, event.Resource)
			}
		case "DELETE":
			if assert.Equal(h.t, tests[2].Resource, event.Resource){
				return nil
			}else{
				return fmt.Errorf("unexpected data: \n%v\n%v", tests[2].Resource, event.Resource)
			}
		}
	}
	return nil
}

func (TestConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (TestConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }