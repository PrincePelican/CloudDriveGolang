package repository

import (
	"cloud-service/entity"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ResourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

func (r *ResourceRepository) GetResourceById(c *gin.Context, id uint64) (entity.ResourceEntity, error) {
	var resource entity.ResourceEntity
	err := r.db.Where("id = ?", id).First(&resource).Error

	return resource, err
}

func (r *ResourceRepository) CreateNewResource(c *gin.Context, entity entity.ResourceEntity) ([]entity.ResourceEntity, error) {
	result := r.db.Create(&entity)

	fmt.Print(result.RowsAffected)

	return nil, nil
}

func (r *ResourceRepository) GetAll(c *gin.Context) ([]entity.ResourceEntity, error) {
	var resources []entity.ResourceEntity
	r.db.Find(&resources)
	return resources, nil
}

func (r *ResourceRepository) ChangeResource(c *gin.Context, entity entity.ResourceEntity, id uint64) error {
	r.db.Save(entity)

	return nil
}

func (r *ResourceRepository) DeleteResource(c *gin.Context, id uint64) error {
	r.db.Delete(&entity.ResourceEntity{}, id)
	return nil
}

func (r *ResourceRepository) GetAllChilds(c *gin.Context, id uint64) (entity.ResourceEntity, error) {
	var entity entity.ResourceEntity
	err := r.db.Model(entity).Preload("ResourceEntity").Where("id = ?", id).Error
	return entity, err
}
