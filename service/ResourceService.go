package service

import (
	dto "cloud-service/DTO"
	"cloud-service/entity"
	"cloud-service/repository"
	"log"
	"time"
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

func (service ResourceService) ChangeResource(dto dto.ResourceDTO, id int64) error {

	dto.ModificationDate = time.Now()
	err := service.resourceRepository.ChangeResource(dto, id)
	if err != nil {
		log.Fatalf("Query error : %s", err)
	}

	return nil
}
