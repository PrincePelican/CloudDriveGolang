package service

import (
	"cloud-service/DTO"
	"cloud-service/converter"
	"cloud-service/entity"
	"cloud-service/repository"
	"cloud-service/validator"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResourceService struct {
	resourceRepository       repository.ResourceRepository
	storageService           StorageService
	downloadDirectorySerivce DownloadDirectorySerivce
}

func NewResourceService(resourceRepository repository.ResourceRepository, storageService StorageService, downloadDirectorySerivce DownloadDirectorySerivce) *ResourceService {
	return &ResourceService{
		resourceRepository:       resourceRepository,
		storageService:           storageService,
		downloadDirectorySerivce: downloadDirectorySerivce,
	}
}

func (service ResourceService) UploadResources(c *gin.Context, resource DTO.FileCreateForm) error {
	if validator.ValidateFileCreateForm(resource.Files, resource.Paths) {
		c.Error(errors.New("Validate: FileCreateForm inncorect"))
	}
	parent, err := service.resourceRepository.GetResourceById(resource.ParentId)
	if err != nil {
		c.Error(err)
	}

	dirStructure, keys := ConvertFromPathsToTreeStructure(resource.Paths, parent.Name)
	resourceStructure := ConvertFromDirStructureToResourceTree(dirStructure)

	service.resourceRepository.CreateNewResource(*resourceStructure)

	for index, file := range resource.Files {
		opened, err := file.Open()
		if err != nil {
			c.Error(err)
		}
		service.storageService.UplodadFileToBucket(c, opened, keys[index])
	}

	return nil
}

func (service ResourceService) GetResourceById(c *gin.Context, id uint64) (string, error) {
	resource, err := service.resourceRepository.GetResourceById(id)
	service.resourceRepository.GetAllChilds(&resource)
	if err != nil {
		c.Error(err)
	}

	startDirectory := (uuid.New()).String()
	service.downloadDirectorySerivce.CreateFilesAndFolderFromResource(c, &resource, startDirectory)
	filepath, err := converter.ZipDir(startDirectory, startDirectory)
	if err != nil {
		c.Error(err)
	}

	return filepath, err
}

func (service ResourceService) GetAll(c *gin.Context) ([]entity.ResourceEntity, error) {
	data, err := service.resourceRepository.GetAll()
	if err != nil {
		c.Error(err)
	}

	return data, nil
}

func (service ResourceService) ChangeResource(c *gin.Context, entity entity.ResourceEntity, id uint64) error {
	_, err := service.resourceRepository.ChangeResource(entity, id)
	if err != nil {
		c.Error(err)
	}

	return nil
}

func (service ResourceService) DeleteResource(c *gin.Context, id uint64) error {
	resource, err := service.resourceRepository.GetResourceById(id)
	if err != nil {
		c.Error(err)
	}
	service.resourceRepository.GetAllChilds(&resource)
	err = service.DeleteResouceRecursive(c, &resource)
	if err != nil {
		c.Error(err)
	}
	return err
}

func (service ResourceService) DeleteResouceRecursive(c *gin.Context, resource *entity.ResourceEntity) error {
	for _, child := range resource.Childs {
		err := service.DeleteResouceRecursive(c, &child)
		if err != nil {
			c.Error(err)
		}
	}

	if len(resource.Key) != 0 {
		service.storageService.DeleteFileFromBucket(c, resource.Key)
	}
	err := service.resourceRepository.DeleteResource(resource)
	if err != nil {
		c.Error(err)
	}
	return err
}
