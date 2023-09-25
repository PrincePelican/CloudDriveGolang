package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"fmt"
	"strings"
)

func FromPathsToTreeStructure(paths []string, DirName string) *DTO.DirectoryNode {
	var structure = DTO.NewDirectoryNode(DirName)
	for _, x := range paths {
		structure.CreateNodes(strings.Split(x, "/"))
	}
	structure.PrintStructure(0)
	return structure
}

func FromDirStructureToResourceTree(dirStructure DTO.DirectoryNode) entity.ResourceEntity {
	mainResource := createEntityFromDirNode(dirStructure)
	createChildOfResource(dirStructure, mainResource)
	fmt.Print(mainResource)
	return *mainResource
}

func createChildOfResource(dirParent DTO.DirectoryNode, parentEntity *entity.ResourceEntity) {
	for _, x := range dirParent.Nodes {
		childEntity := createEntityFromDirNode(*x)
		createChildOfResource(*x, childEntity)
		parentEntity.Childs = append(parentEntity.Childs, *childEntity)
	}
}

func createEntityFromDirNode(dirNode DTO.DirectoryNode) *entity.ResourceEntity {
	return &entity.ResourceEntity{
		Name:         dirNode.Name,
		Key:          "",
		ResourceType: entity.Container,
		Size:         0,
		Childs:       []entity.ResourceEntity{},
	}
}
