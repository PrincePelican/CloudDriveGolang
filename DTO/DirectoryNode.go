package DTO

import "fmt"

type DirectoryNode struct {
	Name  string
	Nodes map[string]*DirectoryNode
}

func NewDirectoryNode(name string) *DirectoryNode {
	return &DirectoryNode{
		Name:  name,
		Nodes: make(map[string]*DirectoryNode),
	}
}

func (p *DirectoryNode) CreateNodes(path []string) {
	if len(path) == 0 {
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
