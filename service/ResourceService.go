package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"cloud-service/repository"
	"os"

	"github.com/gin-gonic/gin"
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

func (service ResourceService) UploadResources(c *gin.Context, resource DTO.FileCreateForm) error {
	var entity entity.ResourceEntity
	parent, err := service.resourceRepository.GetResourceById(c, resource.ParentId)
	if err != nil {
		c.Error(err)
	}

	dirStructure := ConvertFromPathsToTreeStructure(resource.Paths, parent.Name)
	resourceStructure := ConvertFromDirStructureToResourceTree(dirStructure)
	for _, x := range resource.File {

		entity.Key = (uuid.New()).String()
		entity.Name = x.Filename
		entity.Size = x.Size

		file, err := x.Open()
		if err != nil {
			c.Error(err)
		}

		service.storageService.UplodadFileToBucket(c, file, entity.Key)
		service.resourceRepository.CreateNewResource(c, *resourceStructure)
	}

	return nil
}

func (service ResourceService) GetResourceById(c *gin.Context, id uint64) (*os.File, error) {
	resource, err := service.resourceRepository.GetResourceById(c, id)
	if err != nil {
		c.Error(err)
	}

	file := service.storageService.DownloadFileFromBucket(c, resource)

	return file, nil
}

func (service ResourceService) GetAll(c *gin.Context) ([]entity.ResourceEntity, error) {
	data, err := service.resourceRepository.GetAll(c)
	if err != nil {
		c.Error(err)
	}

	return data, nil
}

func (service ResourceService) ChangeResource(c *gin.Context, entity entity.ResourceEntity, id uint64) error {
	err := service.resourceRepository.ChangeResource(c, entity, id)
	if err != nil {
		c.Error(err)
	}

	return nil
}

func (service ResourceService) DeleteResource(c *gin.Context, id uint64) error {
	resource, err := service.resourceRepository.GetResourceById(c, id)
	if err != nil {
		c.Error(err)
	}
	service.storageService.DeleteFileFromBucket(c, resource.Key)
	service.resourceRepository.DeleteResource(c, id)

	return nil
}
