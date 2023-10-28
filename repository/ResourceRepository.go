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

func (r *ResourceRepository) DeleteResource(c *gin.Context, resource *entity.ResourceEntity) error {
	err := r.db.Delete(resource).Error
	return err
}

func (r *ResourceRepository) GetAllChilds(c *gin.Context, resource *entity.ResourceEntity) {
	r.db.Model(resource).Preload("Childs").Where("parent_id = ?", resource.ID).Find(&resource.Childs)
	for i := range resource.Childs {
		r.GetAllChilds(c, &resource.Childs[i])
	}
}
