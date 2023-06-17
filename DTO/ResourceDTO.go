package dto

import "time"

type ResourceDTO struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	ModificationDate time.Time
}
