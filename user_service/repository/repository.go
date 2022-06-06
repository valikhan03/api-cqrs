package repository

import (
	"github.com/jmoiron/sqlx"
	"user_service/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}


func (r *Repository) CreateUser(user models.User) error {
	
}

func (r *Repository) GetUserData(email string) (*models.User, error) {

}