package DTO

import (
	"mime/multipart"
)

type FileCreateForm struct {
	File     []*multipart.FileHeader `form:"file"`
	Paths    []string                `form:"paths"`
	ParentId uint64                  `form:"parentId"`
}
