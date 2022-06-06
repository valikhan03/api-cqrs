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

func NewHandler(usecase Usecase) *Handler{
	return &Handler{usecase: usecase}
}


// Register godoc
// @Summary Register new user
// @Tags user-service
// @Schemes http https
// @Description
// @Accept json
// @Produce json
// @Success 201 {string} status "created"
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
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



// LogIn godoc
// @Summary Log in the user
// @Tags user-service
// @Schemes http https
// @Description
// @Accept json
// @Produce json
// @Success 200 {string} status "authorized"
// @Failure 500 {object} models.ErrorResponse
// @Router /login [post]
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

