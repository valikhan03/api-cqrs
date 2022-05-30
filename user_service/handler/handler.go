package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user_service/models"
)

type Handler struct {
	usecase Usecase
}

type Usecase interface {
	NewUser(user models.User) error
	LogIn(email, password string) (string, error)
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {

	}

	err = h.usecase.NewUser(user)
	if err != nil {

	}

	c.Status(http.StatusCreated)
}

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LogIn(c *gin.Context) {
	var loginData loginData
	err := c.BindJSON(loginData)
	if err != nil{

	}

	token, err := h.usecase.LogIn(loginData.Email, loginData.Password)
	if err != nil{

	}

	c.SetCookie("access-token", token, 100000000, "/api/", "", false, true)
	c.Status(http.StatusOK)
}
