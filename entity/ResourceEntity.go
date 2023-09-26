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
	Name         string           `json:"name"`
	Key          string           `json:"key"`
	ResourceType ResourceType     `json:"resourceType"`
	Size         int64            `json:"size"`
	ParentId     int64            `json:"parentId" gorm:"default:null"`
	Childs       []ResourceEntity `json:"childs" gorm:"foreignKey:ParentId;references:ID"`
}
