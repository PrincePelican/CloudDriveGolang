package controller

import (
	"cloud-service/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceController struct {
	resourceService service.ResourceService
	router          *gin.Engine
}

func NewResourceController(resourceService service.ResourceService, router *gin.Engine) *ResourceController {
	return &ResourceController{
		resourceService: resourceService,
		router:          router,
	}
}

func (ctr ResourceController) InitRoutes() {
	ctr.router.GET("/resources/all", ctr.getAll)
}

func (ctr ResourceController) getAll(c *gin.Context) {
	data, err := ctr.resourceService.GetAll()
	if err != nil {
		log.Fatalf("error %s", err)
	}
	c.IndentedJSON(http.StatusOK, data)
}
