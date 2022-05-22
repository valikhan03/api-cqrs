package repository

import (
	"gorm.io/gorm"

	"resource_admin/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateResource(resource *models.Resource) error{
	err := r.db.Create(resource).Error
	return err
}

func (r *Repository) UpdateResource(resource *models.Resource) error{
	err := r.db.Model(&resource).Where("id=?", resource.ID).Updates(&resource).Error
	return err
}

func (r *Repository) DeleteResource(id string) error{
	err := r.db.Delete(&models.Resource{}, id).Error
	return err
}