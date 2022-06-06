package repository

import (
	"context"
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


func (r *Repository) CreateUser(ctx context.Context, user models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES $1, $2, $3"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}

func (r *Repository) GetUserData(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT name, email, password FROM users WHERE email=$1"
	err := r.db.GetContext(ctx, &user, query, email)
	return &user, err
}