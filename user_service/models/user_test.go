package models

import(
	"os"
	"testing"

	"user_service/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	db, err := repository.InitDB()
	assert.NoError(t, err)

	err = db.AutoMigrate(&User{})
	assert.NoError(t, err)

	user := User{
		Name: "TestName",
		Email: "test@email.com",
		Password: os.Getenv("passwordHash"),
	}
	
	err = user.CreateRecord(db)
	assert.NoError(t, err)

	db.Where("email=?", user.Email).Find(&userResult)

	db.Unscoped().Delete(&user)

	assert.Equal(t, "TestName", userResult.Name)
	assert.Equal(t, "test@email.com", userResult.Email)
}

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("secret")
	assert.NoError(t, err)
}

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword()
	assert.NoError(t, err)
	os.Setenv("passwordHash", user.Password)
}