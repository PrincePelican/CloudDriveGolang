package service

import (
	"cloud-service/entity"
	"cloud-service/repository"
	"log"
)

type ResourceService struct {
	resourceRepository repository.ResourceRepository
}

func NewResourceService(resourceRepository repository.ResourceRepository) *ResourceService {
	return &ResourceService{
		resourceRepository: resourceRepository,
	}
}

func (service ResourceService) CreateResource(entity entity.ResourceEntity) error {
	service.resourceRepository.CreateNewResource(entity)

	return nil
}

func (service ResourceService) GetAll() ([]entity.ResourceEntity, error) {
	data, err := service.resourceRepository.GetAll()
	if err != nil {
		log.Fatalf("Service error %s", err)
	}

	return data, nil
}

func (service ResourceService) ChangeResource(entity entity.ResourceEntity, id uint64) error {
	err := service.resourceRepository.ChangeResource(entity, id)
	if err != nil {
		log.Fatalf("Service error : %s", err)
	}

	return nil
}

func (service ResourceService) DeleteResource(id uint64) error {
	service.resourceRepository.DeleteResource(id)
	return nil
}
