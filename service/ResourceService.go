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

func (service ResourceService) CreateResource() error {
	return nil
}

func (service ResourceService) GetAll() ([]entity.ResourceEntity, error) {
	data, err := service.resourceRepository.GetAll()
	if err != nil {
		log.Fatalf("error %s", err)
	}

	return data, nil
}
