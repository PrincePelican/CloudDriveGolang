package DTO

import (
	"cloud-service/entity"
	"fmt"

	"github.com/google/uuid"
)

type DirectoryNode struct {
	Name         string
	ResourceType entity.ResourceType
	Nodes        map[string]*DirectoryNode
	Key          string
}

func NewDirectoryNode(name string) *DirectoryNode {
	return &DirectoryNode{
		Name:  name,
		Nodes: make(map[string]*DirectoryNode),
	}
}

func NewFileNode(name string, keys *[]string) *DirectoryNode {
	key := (uuid.New()).String()
	*keys = append(*keys, key)
	return &DirectoryNode{
		Name:         name,
		ResourceType: entity.File,
		Key:          key,
	}
}

func (p *DirectoryNode) CreateNodes(path []string, keys *[]string) {
	if len(path) == 1 {
		newNode := NewFileNode(path[0], keys)
		p.Nodes[path[0]] = newNode
		return
	}

	if p.Nodes[path[0]] == nil {
		newNode := NewDirectoryNode(path[0])
		p.Nodes[path[0]] = newNode
	}

	p.Nodes[path[0]].CreateNodes(path[1:], keys)
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
