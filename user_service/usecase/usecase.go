package usecase

import(
	"context"
	"time"

	"user_service/models"
)

type UseCase struct{
	repository Repository
	jwt *models.JWTWrapper
}

type Repository interface{
	CreateUser(ctx context.Context, user models.User) error
	GetUserData(ctx context.Context, email string) (*models.User, error)
}

func NewUseCase(repository Repository) *UseCase{
	return &UseCase{repository: repository}
}

func(uc *UseCase) NewUser(user models.User) error {
	err := user.HashPassword()
	if err != nil{
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	return uc.repository.CreateUser(ctx, user)
}

func(uc *UseCase) LogIn(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	user, err := uc.repository.GetUserData(ctx, email)
	if err != nil{
		return "", err
	}

	err = user.CheckPassword(password)
	if err != nil{
		return "", err
	}

	token, err := uc.jwt.GenerateToken(email)
	if err != nil{
		return "", err
	}

	return token, nil
}