package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"strings"

	"github.com/google/uuid"
)

func ConvertFromPathsToTreeStructure(paths []string, DirName string) *DTO.DirectoryNode {
	var structure = DTO.NewDirectoryNode(DirName)
	for _, x := range paths {
		structure.CreateNodes(strings.Split(x, "/"))
	}
	return structure
}

func ConvertFromDirStructureToResourceTree(dirStructure *DTO.DirectoryNode) *entity.ResourceEntity {
	mainResource := createEntityFromDirNode(*dirStructure)
	createChildFromResource(*dirStructure, mainResource)
	return mainResource
}

func createChildFromResource(dirParent DTO.DirectoryNode, parentEntity *entity.ResourceEntity) {
	for _, x := range dirParent.Nodes {
		childEntity := createEntityFromDirNode(*x)
		createChildFromResource(*x, childEntity)
		parentEntity.Childs = append(parentEntity.Childs, *childEntity)
	}
}

func createEntityFromDirNode(dirNode DTO.DirectoryNode) *entity.ResourceEntity {
	key := ""
	if dirNode.ResourceType == entity.File {
		key = (uuid.New()).String()
	}
	return &entity.ResourceEntity{
		Name:         dirNode.Name,
		Key:          key,
		ResourceType: entity.Container,
		Size:         0,
		Childs:       []entity.ResourceEntity{},
	}
}
