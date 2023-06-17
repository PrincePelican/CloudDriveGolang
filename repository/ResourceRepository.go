package repository

import (
	"cloud-service/entity"
	"database/sql"
	"log"
)

type ResourceRepository struct {
	db *sql.DB
}

func NewResourceRepository(db *sql.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

func (r *ResourceRepository) GetAll() ([]entity.ResourceEntity, error) {
	data, err := r.db.Query("SELECT * FROM RESOURCES")
	if err != nil {
		log.Fatalf("Query error : %s", err)
	}

	var resources []entity.ResourceEntity

	for data.Next() {
		var rsc entity.ResourceEntity

		if err := data.Scan(&rsc.ID, &rsc.Name, &rsc.Path,
			&rsc.ResourceType, &rsc.Size, &rsc.ModificationDate, &rsc.ParentId); err != nil {
			log.Fatalf("Scan data error : %s", err)
		}

		resources = append(resources, rsc)
	}

	return resources, nil
}
