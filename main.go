package main

import (
	"cloud-service/controller"
	"cloud-service/dbproperties"
	"cloud-service/entity"
	"cloud-service/repository"
	"cloud-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getRsc(c *gin.Context, rsc []entity.ResourceEntity) {
	c.IndentedJSON(http.StatusOK, rsc)
}

func main() {
	db := dbproperties.Connection()
	router := gin.Default()
	rscR := repository.NewResourceRepository(db)
	rscS := service.NewResourceService(*rscR)
	rscC := controller.NewResourceController(*rscS, router)

	rscC.InitRoutes()

	router.Run("localhost:8080")
}
