package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"strings"
)

func FromPathsToTreeStructure(paths []string) *DTO.DirectoryNode {
	var structure = DTO.NewDirectoryNode()
	for _, x := range paths {
		structure.CreateNodes(strings.Split(x, "/"))
	}
	structure.PrintStructure(0)
	return structure
}

func CompareStructureToResources(dirStructure DTO.DirectoryNode, resourceEntity entity.ResourceEntity) {

}
