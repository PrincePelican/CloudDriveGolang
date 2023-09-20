package DTO

import (
	"mime/multipart"
)

type FilesStructure struct {
	DirectoryName string           `json:"name"`
	Files         []string         `json:"files"`
	Directory     []FilesStructure `json:"folders"`
}

type FileCreateForm struct {
	File []*multipart.FileHeader `form:"file"`
}
