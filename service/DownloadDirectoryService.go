package service

import (
	"cloud-service/entity"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DownloadDirectorySerivce struct {
	storageService StorageService
	tmp            string
}

func NewDownloadDirectorySerivce(storageService StorageService) *DownloadDirectorySerivce {
	var tmp string
	if env := os.Getenv("TmpFolder"); env != "" {
		tmp = env
	}
	return &DownloadDirectorySerivce{
		storageService: storageService,
		tmp:            tmp,
	}
}

func (s *DownloadDirectorySerivce) CreateFilesAndFolderFromResource(c *gin.Context, resource *entity.ResourceEntity, path string) error {
	if len(resource.Key) != 0 {
		s.storageService.DownloadFileFromBucket(c, *resource, filepath.Join(s.tmp, path))
	} else {
		path = filepath.Join(path, resource.Name)
		s.createFolderFromResource(c, resource, &path)
	}

	return nil
}

func (s *DownloadDirectorySerivce) createFolderFromResource(c *gin.Context, resource *entity.ResourceEntity, path *string) error {
	err := os.MkdirAll(filepath.Join(s.tmp, *path), os.ModePerm)
	if err != nil {
		c.Error(err)
	}

	for _, child := range resource.Childs {
		err := s.CreateFilesAndFolderFromResource(c, &child, *path)
		if err != nil {
			c.Error(err)
		}
	}

	return nil
}
