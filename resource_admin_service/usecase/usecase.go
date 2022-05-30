package usecase

import (
	"log"
	"resource_admin_service/models"
	"time"

	"github.com/google/uuid"
)

type UseCase struct {
	logger     *log.Logger
	repository RepositoryInterface
	eventChan  chan<- *models.Event
}

func NewUseCase(repository RepositoryInterface, eventChan chan *models.Event, logger *log.Logger) *UseCase {
	return &UseCase{
		repository: repository,
		logger: logger,
		eventChan: eventChan,

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
		CreatedAt: time.Now().UTC(),
	}

	err := uc.repository.CreateResource(&resource)
	if err != nil {
		return err
	}
	event := models.NewEvent("CREATE", resource)

	uc.eventChan <- event
	return nil
}

func (uc *UseCase) UpdateResource(resource *models.Resource) error {
	err := uc.repository.UpdateResource(resource)
	if err != nil {
		return err
	}
	event := models.NewEvent("UPDATE", *resource)
	
	uc.eventChan <- event
	return nil
}

func (uc *UseCase) DeleteResource(id string) error {
	err := uc.repository.DeleteResource(id)
	if err != nil {
		return err
	}
	resource := models.Resource{ID: id}
	event := models.NewEvent("DELETE", resource)
	uc.eventChan <- event

	return nil
}
