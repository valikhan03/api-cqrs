package repository

import (
	"github.com/jmoiron/sqlx"

	"resource_admin_service/models"
)

type RepositorySQLX struct {
	db *sqlx.DB
}

func NewRepositorySQLX(db *sqlx.DB) *RepositorySQLX {
	return &RepositorySQLX{db: db}
}

func (r *RepositorySQLX) CreateResource(resource *models.Resource) error {
	query := "INSERT INTO resources (id, title, author, content, created_at, tags) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, resource.ID, resource.Title, resource.Author, resource.Content, resource.CreatedAt, resource.Tags)
	return err
}

func (r *RepositorySQLX) UpdateResource(resource *models.Resource) error {
	query := "UPDATE resources (title, author, content, created_at, tags) SET ($1, $2, $3, $4, $5) WHERE id=$6"
	_, err := r.db.Exec(query, resource.Title, resource.Author, resource.Content, resource.CreatedAt, resource.Tags, resource.ID)
	return err
}

func (r *RepositorySQLX) DeleteResource(id string) error {
	query := "DELETE FROM resources WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}
