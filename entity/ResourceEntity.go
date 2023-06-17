package entity

import "time"

type ResourceType uint

const (
	File      ResourceType = 0
	Container ResourceType = 1
)

type ResourceEntity struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	Path             string       `json:"path"`
	ResourceType     ResourceType `json:"resourceType"`
	Size             int64        `json:"size"`
	ModificationDate time.Time    `json:"modificationDate"`
	ParentId         int32        `json:"parentId"`
}
