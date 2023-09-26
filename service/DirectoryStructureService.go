package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"strings"
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
	return &entity.ResourceEntity{
		Name:         dirNode.Name,
		Key:          "",
		ResourceType: entity.Container,
		Size:         0,
		Childs:       []entity.ResourceEntity{},
	}
}
