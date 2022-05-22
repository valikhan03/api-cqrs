package usecase

import (
	"log"
	"resource_admin/models"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type UseCase struct {
	logger *log.Logger
	repository RepositoryInterface
	producer   sarama.SyncProducer
}

func NewUseCase(rep RepositoryInterface, prod sarama.SyncProducer) *UseCase {
	return &UseCase{
		repository: rep,
		producer: prod,
	}
}

type RepositoryInterface interface {
	CreateResource(resource *models.Resource) error
	UpdateResource(resource *models.Resource) error
	DeleteResource(id string) error
}

func (uc *UseCase) CreateResource(newresource *models.NewResource) error {
	resource := models.Resource{
		ID:        uuid.NewString(),
		Title:     newresource.Title,
		Author:    newresource.Author,
		Content:   newresource.Content,
		CreatedAt: time.Now(),
	}

	err := uc.repository.CreateResource(&resource)
	if err != nil{
		return err
	}
	event, err := models.NewEvent("CREATE", resource).Encode()
	if err != nil{
		uc.logger.Println(err)
	}
	uc.sendEvent("CREATE", event)
	return nil
}

func (uc *UseCase) UpdateResource(resource *models.Resource) error {
	err :=  uc.repository.UpdateResource(resource)
	if err != nil{
		return err
	}
	event, err := models.NewEvent("UPDATE", *resource).Encode()
	if err != nil{
		uc.logger.Println(err)
	}
	uc.sendEvent("UPDATE", event)
	return nil
}

func (uc *UseCase) DeleteResource(id string) error {
	err :=  uc.repository.DeleteResource(id)
	if err != nil{
		return err
	}
	resource := models.Resource{ID: id}
	event, err := models.NewEvent("CREATE", resource).Encode()
	if err != nil{
		uc.logger.Println(err)
	}
	uc.sendEvent("CREATE", event)
	return nil
}

func (uc *UseCase) sendEvent(topic string, event []byte) {
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(event),
	}
	uc.producer.SendMessage(&msg)
}
