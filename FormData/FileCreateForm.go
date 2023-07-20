package formdata

import (
	"cloud-service/entity"
	"mime/multipart"
)

type FileCreateForm struct {
	File     *multipart.FileHeader `form:"file"`
	Resource entity.ResourceEntity `form:"resource"`
}
