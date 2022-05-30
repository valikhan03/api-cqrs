package usecase

import(
	"user_service/models"
)

type UseCase struct{
	repository Repository
	jwt *models.JWTWrapper
}

type Repository interface{
	CreateUser(user models.User) error
	GetUserData(email string) (*models.User, error)
}

func NewUseCase() {

}

func(uc *UseCase) NewUser(user models.User) error {
	user.HashPassword()
	return uc.repository.CreateUser(user)
}

func(uc *UseCase) LogIn(email, password string) (string, error) {
	user, err := uc.repository.GetUserData(email)
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