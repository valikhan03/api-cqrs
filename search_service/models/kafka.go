package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct{
	logger *log.Logger
	handler *ConsumerGroupHandler
}

func NewKafkaConsumer(logger *log.Logger, handler *ConsumerGroupHandler) *KafkaConsumer{
	return &KafkaConsumer{
		logger: logger,
		handler: handler,
	}
}

func (kc *KafkaConsumer) Consume() {
	configs := sarama.NewConfig()
	configs.Consumer.Return.Errors = true
	configs.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup([]string{os.Getenv("KAFKA_CONSUMER_GROUP_ADDR")}, os.Getenv("KAFKA_CONSUMER_GROUP_ID"), configs)
	if err != nil{
		kc.logger.Fatal(err)
	}

	go func(){
		for err := range consumerGroup.Errors(){
			kc.logger.Println(err)
		}
	}()

	for {
		err := consumerGroup.Consume(context.Background(), []string{}, kc.handler)
		if err != nil{
			kc.logger.Println(err)
		}
	}
}

type ConsumerGroupHandler struct{
	logger *log.Logger
	elatic *Elastic
}

func NewConsumerGroupHandler(addr []string, index string, timeout int64, logger *log.Logger) *ConsumerGroupHandler{
	return &ConsumerGroupHandler{
		elatic: NewElastic(addr, index, timeout),
		logger: logger,
	}
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.logger.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		
		event, err := DecodeEvent(msg.Value)
		if err != nil{
			h.logger.Println(fmt.Errorf("unable to decode event: %v", err))
			continue
		}

		switch event.EventType{
		case CreateEvent:
			err := h.elatic.CreateResource(event.Resource)
			if err != nil{
				h.logger.Printf("unable to create event: %v\n", err)
			}
		case UpdateEvent:
			err := h.elatic.UpdateResource(event.Resource)
			if err != nil{
				h.logger.Printf("unable to update event: %v\n", err)
			}
		case DeleteEvent:
			err := h.elatic.DeleteResource(event.Resource.ID)
			if err != nil{
				h.logger.Printf("unable to delete event: %v\n", err)
			}
		}
		
		session.MarkMessage(msg, "done")
	}
	return nil
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }