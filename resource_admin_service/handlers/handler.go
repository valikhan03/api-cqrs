package handlers

import (
	"errors"
	"log"
	"net/http"
	"resource_admin/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase Usecase
	logger  *log.Logger
}

func NewHandler(uc Usecase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

type Usecase interface {
	CreateResource(resource *models.NewResource) error
	UpdateResource(resource *models.Resource) error
	DeleteResource(id string) error
}

// CreateResource godoc
// @Summary Create Resource Handler
// @Tags resource-service
// @Schemes
// @Description The endpoint for creating new Resource-objects in service DBs
// @Accept json
// @Produce json
// @Success 201 {string} status "created"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router / [post]
func (h *Handler) CreateResourceHandler(c *gin.Context) {
	var resource models.NewResource
	err := c.BindJSON(&resource)
	if err != nil {
		//log
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid format of data"))
		return
	}

	err = h.usecase.CreateResource(&resource)
	if err != nil {
		//log
		h.logger.Printf("Create resource: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateResource godoc
// @Summary Update Resource Handler
// @Tags resource-service
// @Schemes http https
// @Description The endpoint for updating existing resources
// @Accept json
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /{id} [post]
func (h *Handler) UpdateResourceHandler(c *gin.Context) {
	c.Param("id")
	var resource models.Resource
	err := c.BindJSON(&resource)
	if err != nil {
		//log
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid format of data"))
		return
	}

	err = h.usecase.UpdateResource(&resource)
	if err != nil {
		h.logger.Printf("Update resource: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// DeleteResource godoc
// @Summary Delete Resource Handler
// @Tags resource-service
// @Schemes http https
// @Description The endpoint to delete Resource from DB
// @Accept plain
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /{id} [delete]
func (h *Handler) DeleteResourceHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		//log
		c.AbortWithError(http.StatusBadRequest, errors.New("empty id"))
		return
	}
	err := h.usecase.DeleteResource(id)
	if err != nil {
		log.Printf("Delete Resource: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
