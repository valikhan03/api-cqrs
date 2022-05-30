package models

import (
	"log"

	"github.com/Shopify/sarama"
)

type Kafka struct{
	producer sarama.SyncProducer
	eventChan <- chan *Event
	logger *log.Logger
}

func NewKafka(producer sarama.SyncProducer, eventChan <- chan *Event, logger *log.Logger) *Kafka{
	return &Kafka{
		producer: producer,
		eventChan: eventChan,
		logger: logger,
	}
}

// goroutine
func (k *Kafka) ExecuteEvents(){
	for{
		event := <- k.eventChan
		eventData, err := event.Encode()
		if err != nil{
			k.logger.Println(err)
			continue
		}
		msg := sarama.ProducerMessage{
			Topic: event.EventType,
			Value: sarama.ByteEncoder(eventData),
		}
		_,_, err = k.producer.SendMessage(&msg)
		k.logger.Println(err)	
	}
}