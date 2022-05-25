package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"search_service/models"
)

type Handler struct{
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Service interface{
	GetResourceByID(id string) (*models.Resource, error)
	SearchResourcesByFilter(filter map[string]interface{}) ([]*models.Resource, error)
}

// GetResourceByID godoc
// @Summary Get resource by id
// @Tags search-service
// @Schemes http https
// @Description returns models.Resource object in json format which matches given id
// @Accept plain
// @Produce json
// @Param id path string true "Resource ID"
// @Success 200 {object} models.Resource
// @Failure 404 {object} models.ErrorResponse
// @Router /{id} [get]
func (h *Handler) GetResourceByID(c *gin.Context) {
	id := c.Param("id")
	resource, err := h.service.GetResourceByID(id)
	if err != nil{

	}

	c.JSON(http.StatusOK, resource)
}

// SearchResourceByFilter godoc
// @Summary get resources by filter
// @Tags search-service
// @Schemes http https
// @Description returns models.Resource array of json-objects which match filter params
// @Accept plain
// @Produce json
// @Success 200 {object} []models.Resource
// @Failure 404 {object} models.ErrorResponse
// @Router /search [get]
func (h *Handler) SearchResourcesByFilter(c *gin.Context) {
	/*
	var filter models.Filter
	err := c.BindJSON(&filter)
	if err != nil{
		return
	}
	*/

	filters := make(map[string]interface{})
	_, ok := c.GetQuery("title")
	if ok{
		filters["title"] = c.Query("title")
	}
	_, ok = c.GetQueryArray("author")
	if ok{
		filters["author"] = c.QueryArray("author")
	}
	_, ok = c.GetQuery("content")
	if ok{
		filters["content"] = c.Query("content")
	}
	_, ok = c.GetQueryArray("tags")
	if ok{
		filters["tags"] = c.QueryArray("tags")
	}

	resources, err := h.service.SearchResourcesByFilter(filters)
	if err != nil{
		return
	}

	c.SecureJSON(http.StatusOK, resources)
}