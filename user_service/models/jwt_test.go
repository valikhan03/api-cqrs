package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	jwtWrapper := JWTWrapper{
		SecretKey:      "test-secret-key",
		Issuer:         "auth-service",
		ExpirationTime: 2,
	}

	generatedToken, err := jwtWrapper.GenerateToken("test@mail.com")
	assert.NoError(t, err)

	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	encodedToken := os.Getenv("testToken")

	jwtWrapper := JWTWrapper{
		SecretKey: "test-secret-key",
		Issuer:    "auth-service",
	}

	claims, err := jwtWrapper.ValidateToken(encodedToken)
	assert.NoError(t, err)

	assert.Equal(t, "auth-service", claims.Issuer)
	assert.Equal(t, "test@mail.com", claims.Email)
}
