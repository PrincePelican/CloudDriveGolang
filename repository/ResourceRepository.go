package repository

import (
	"cloud-service/entity"

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

func (r *ResourceRepository) GetResourceById(id uint64) (entity.ResourceEntity, error) {
	var resource entity.ResourceEntity
	err := r.db.Where("id = ?", id).First(&resource).Error

	return resource, err
}

func (r *ResourceRepository) CreateNewResource(entity entity.ResourceEntity) (entity.ResourceEntity, error) {
	result := r.db.Create(&entity)
	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}

func (r *ResourceRepository) GetAll() ([]entity.ResourceEntity, error) {
	var resources []entity.ResourceEntity
	r.db.Find(&resources)
	return resources, nil
}

func (r *ResourceRepository) ChangeResource(entity entity.ResourceEntity, id uint64) (entity.ResourceEntity, error) {
	result := r.db.Save(entity)
	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}

func (r *ResourceRepository) DeleteResource(resource *entity.ResourceEntity) error {
	err := r.db.Delete(resource).Error
	return err
}

func (r *ResourceRepository) GetAllChilds(resource *entity.ResourceEntity) {
	r.db.Model(resource).Preload("Childs").Where("parent_id = ?", resource.ID).Find(&resource.Childs)
	for i := range resource.Childs {
		r.GetAllChilds(&resource.Childs[i])
	}
}
