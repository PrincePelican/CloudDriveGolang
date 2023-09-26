package service

import (
	"cloud-service/entity"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

type StorageService struct {
	session *session.Session
}

func NewStorageService() *StorageService {
	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-north-1")}))
	return &StorageService{
		session: session,
	}
}

func (s *StorageService) UplodadFileToBucket(c *gin.Context, file io.Reader, key string) {

	uploader := s3manager.NewUploader(s.session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("cloudgodrivebucket"),
		Body:   file,
		Key:    aws.String(key),
	})

	if err != nil {
		c.Error(err)
	}

	fmt.Print(result)
}

func (s *StorageService) DownloadFileFromBucket(c *gin.Context, resource entity.ResourceEntity) *os.File {
	downloader := s3manager.NewDownloader(s.session)

	file, err := os.Create(resource.Name)
	if err != nil {
		c.Error(err)
	}

	defer file.Close()

	result, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("cloudgodrivebucket"),
			Key:    aws.String(resource.Key),
		})
	if err != nil {
		c.Error(err)
	}
	fmt.Println("Downloaded bytes", result)

	return file
}

func (s *StorageService) DeleteFileFromBucket(c *gin.Context, key string) {
	svc := s3.New(s.session)

	result, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("cloudgodrivebucket"),
		Key:    aws.String(key),
	})

	if err != nil {
		c.Error(err)
	}

	fmt.Print(result)

}
