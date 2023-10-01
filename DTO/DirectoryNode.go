package DTO

import (
	"cloud-service/entity"
	"fmt"
)

type DirectoryNode struct {
	Name         string
	ResourceType entity.ResourceType
	Nodes        map[string]*DirectoryNode
}

func NewDirectoryNode(name string) *DirectoryNode {
	return &DirectoryNode{
		Name:  name,
		Nodes: make(map[string]*DirectoryNode),
	}
}

func NewFileNode(name string) *DirectoryNode {
	return &DirectoryNode{
		Name:         name,
		ResourceType: entity.File,
	}
}

func (p *DirectoryNode) CreateNodes(path []string) {
	if len(path) == 1 {
		newNode := NewFileNode(path[0])
		p.Nodes[path[0]] = newNode
		return
	}

	if p.Nodes[path[0]] == nil {
		newNode := NewDirectoryNode(path[0])
		p.Nodes[path[0]] = newNode
	}

	p.Nodes[path[0]].CreateNodes(path[1:])
}

func (p *DirectoryNode) PrintStructure(deep int) {
	fmt.Println()
	for element := range p.Nodes {
		for i := 0; i < deep; i++ {
			fmt.Printf(" ")
		}

		fmt.Printf(element)
		p.Nodes[element].PrintStructure(deep + 1)
	}
}
