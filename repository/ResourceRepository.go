package repository

import (
	dto "cloud-service/DTO"
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

func (r *ResourceRepository) ChangeResource(dto dto.ResourceDTO, id int64) error {
	_, err := r.db.Query("UPDATE RESOURCES SET name = $1, path = $2, modification_date = $3 WHERE id = $4", dto.Name, dto.Path, dto.ModificationDate, id)
	if err != nil {
		log.Fatalf("Query error : %s", err)
	}
	return nil
}
