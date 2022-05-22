package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"search_service/models"
)

type Handler struct{
	service Service
}

type Service interface{
	GetResourceByID(id string) (*models.Resource, error)
	SearchResourcesByFilter(filter models.Filter) ([]*models.Resource, error)
}

// GetResourceByID godoc
// @Description returns models.Resource object in json format which matches given id
// @Summary get resource by id
// @Tags Resource
// @Accept param id string
// @Produce json
// @Success 200
// @Failure 404
// @Router /api/v1/{id}
func (h *Handler) GetResourceByID(c *gin.Context) {
	id := c.Param("id")
	resource, err := h.service.GetResourceByID(id)
	if err != nil{

	}

	c.JSON(http.StatusOK, resource)
}


func (h *Handler) SearchResourcesByFilter(c *gin.Context) {
	var filter models.Filter
	err := c.BindJSON(&filter)
	if err != nil{
		return
	}

	resources, err := h.service.SearchResourcesByFilter(filter)
	if err != nil{
		return
	}

	c.SecureJSON(http.StatusOK, resources)
}