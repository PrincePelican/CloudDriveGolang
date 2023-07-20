package main

import (
	"cloud-service/controller"
	"cloud-service/dbproperties"
	"cloud-service/repository"
	"cloud-service/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := dbproperties.Connection()
	router := gin.Default()
	s3Service := service.NewStorageService()
	rscR := repository.NewResourceRepository(db)
	rscS := service.NewResourceService(*rscR, *s3Service)
	rscC := controller.NewResourceController(*rscS, router)
	dbproperties.InitTables(db)
	rscC.InitRoutes()

	router.Run("0.0.0.0:8080")
}
