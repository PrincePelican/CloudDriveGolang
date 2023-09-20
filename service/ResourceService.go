package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"cloud-service/repository"
	"log"
	"os"

	"github.com/google/uuid"
)

type ResourceService struct {
	resourceRepository repository.ResourceRepository
	storageService     StorageService
}

func NewResourceService(resourceRepository repository.ResourceRepository, storageService StorageService) *ResourceService {
	return &ResourceService{
		resourceRepository: resourceRepository,
		storageService:     storageService,
	}
}

func (service ResourceService) UploadResources(resource DTO.FileCreateForm) error {
	var entity entity.ResourceEntity

	for _, x := range resource.File {

		entity.Key = (uuid.New()).String()
		entity.Name = x.Filename
		entity.Size = x.Size

		file, err := x.Open()
		if err != nil {
			log.Fatalf("File open error %s", err)
		}

		service.storageService.UplodadFileToBucket(file, entity.Key)
		service.resourceRepository.CreateNewResource(entity)
	}

	return nil
}

func (service ResourceService) GetResourceById(id uint64) (*os.File, error) {
	resource := service.resourceRepository.GetResourceById(id)

	file := service.storageService.DownloadFileFromBucket(resource)

	return file, nil
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
	resource := service.resourceRepository.GetResourceById(id)
	service.storageService.DeleteFileFromBucket(resource.Key)
	service.resourceRepository.DeleteResource(id)

	return nil
}
