package models_test

import(
	"os"
	"testing"

	"user_service/models"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := models.User{
		Password: "secret",
	}

	err := user.HashPassword()
	assert.NoError(t, err)
	os.Setenv("passwordHash", user.Password)
}

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := models.User{
		Password: hash,
	}

	err := user.CheckPassword("secret")
	assert.NoError(t, err)
}

