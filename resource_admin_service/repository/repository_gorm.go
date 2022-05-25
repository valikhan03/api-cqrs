package repository

import (
	"gorm.io/gorm"

	"resource_admin_service/models"
)

type RepositoryGORM struct {
	db *gorm.DB
}

func NewRepositoryGORM(db *gorm.DB) *RepositoryGORM {
	return &RepositoryGORM{db: db}
}

func (r *RepositoryGORM) CreateResource(resource *models.Resource) error {
	err := r.db.Create(resource).Error
	return err
}

func (r *RepositoryGORM) UpdateResource(resource *models.Resource) error {
	err := r.db.Model(&resource).Where("id=?", resource.ID).Updates(&resource).Error
	return err
}

func (r *RepositoryGORM) DeleteResource(id string) error {
	err := r.db.Delete(&models.Resource{}, id).Error
	return err
}
