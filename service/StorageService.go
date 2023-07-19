package service

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func (s *StorageService) UplodadFileToBucket(file io.Reader, key string) {

	uploader := s3manager.NewUploader(s.session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("cloudgodrivebucket"),
		Body:   file,
		Key:    aws.String(key),
	})

	if err != nil {
		log.Fatalf("Upload error %s", err)
	}

	fmt.Print(result)

}

func (s *StorageService) DownloadFileFromBucket(key string) {
	downloader := s3manager.NewDownloader(s.session)

	file, err := os.Create("pobrany.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	result, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String("cloudgodrivebucket"),
			Key:    aws.String(key),
		})
	if err != nil {
		log.Fatalf("Download error %s", err)
	}

	fmt.Println("Downloaded", file.Name(), result, "bytes")
}

func (s *StorageService) DeleteFileFromBucket(key string) {
	svc := s3.New(s.session)

	result, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("cloudgodrivebucket"),
		Key:    aws.String(key),
	})

	if err != nil {
		log.Fatalf("Delete error %s", err)
	}

	fmt.Print(result)

}
