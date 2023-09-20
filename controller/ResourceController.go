package controller

import (
	"cloud-service/DTO"
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
	ctr.router.GET("/resources/all", ctr.getAllResource)
	ctr.router.GET("/resources/:id", ctr.getResourceById)
	ctr.router.PUT("/resources/change/:id", ctr.PutResourceChange)
	ctr.router.POST("/resources/create", ctr.createNewResource)
	ctr.router.DELETE("/resources/delete/:id", ctr.DeleteById)
}

func (ctr ResourceController) createNewResource(c *gin.Context) {
	var newResource DTO.FileCreateForm
	if err := c.ShouldBind(&newResource); err != nil {
		log.Fatalf("error bind %s", err)
	}

	ctr.resourceService.UploadResources(newResource)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "File Uploaded"})
}

func (ctr ResourceController) getAllResource(c *gin.Context) {
	data, err := ctr.resourceService.GetAll()
	if err != nil {
		log.Fatalf("error %s", err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (ctr ResourceController) getResourceById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error id %s", err)
	}
	file, err := ctr.resourceService.GetResourceById(id)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	c.Header("Content-Disposition", "attachment; filename="+file.Name())
	c.Header("Content-Type", "application/octet-stream")
	c.File(file.Name())
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

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "File changed"})
}

func (ctr ResourceController) DeleteById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalf("error id %s", err)
	}

	ctr.resourceService.DeleteResource(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "File deleted"})
}
