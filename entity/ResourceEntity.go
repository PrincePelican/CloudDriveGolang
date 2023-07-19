package entity

import (
	"gorm.io/gorm"
)

type ResourceType uint

const (
	File      ResourceType = 0
	Container ResourceType = 1
)

func (ResourceEntity) TableName() string {
	return "resources"
}

type ResourceEntity struct {
	gorm.Model
	Name         string       `json:"name"`
	Path         string       `json:"path" gorm:"unique"`
	ResourceType ResourceType `json:"resourceType"`
	Size         int64        `json:"size"`
	ParentId     int32        `json:"parentId"`
}
