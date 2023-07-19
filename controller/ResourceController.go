package controller

import (
	"cloud-service/entity"
	"cloud-service/service"
	"log"
	"net/http"
	"strconv"

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
	ctr.router.PUT("/resources/change/:id", ctr.PutResourceChange)
	ctr.router.POST("/resources/create", ctr.createNewResource)
	ctr.router.DELETE("/resources/delete/:id", ctr.DeleteById)
}

func (ctr ResourceController) createNewResource(c *gin.Context) {
	var newResource entity.ResourceEntity
	if err := c.BindJSON(&newResource); err != nil {
		log.Fatalf("error bind %s", err)
	}

	ctr.resourceService.CreateResource(newResource)
}

func (ctr ResourceController) getAll(c *gin.Context) {
	data, err := ctr.resourceService.GetAll()
	if err != nil {
		log.Fatalf("error %s", err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (ctr ResourceController) PutResourceChange(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error id %s", err)
	}
	var changedResource entity.ResourceEntity
	if err := c.BindJSON(&changedResource); err != nil {
		log.Fatalf("error bind %s", err)
	}
	err = ctr.resourceService.ChangeResource(changedResource, id)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Data changed"})
}

func (ctr ResourceController) DeleteById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error id %s", err)
	}

	ctr.resourceService.DeleteResource(id)
}
