package DTO

import "fmt"

type DirectoryNode struct {
	nodes map[string]*DirectoryNode
}

func NewDirectoryNode() *DirectoryNode {
	return &DirectoryNode{
		nodes: make(map[string]*DirectoryNode),
	}
}

func (p *DirectoryNode) CreateNodes(path []string) {
	if len(path) == 0 {
		return
	}

	if p.nodes[path[0]] == nil {
		newNode := NewDirectoryNode()
		p.nodes[path[0]] = newNode
	}

	p.nodes[path[0]].CreateNodes(path[1:])
}

func (p *DirectoryNode) PrintStructure(deep int) {
	fmt.Println()
	for element := range p.nodes {
		for i := 0; i < deep; i++ {
			fmt.Printf(" ")
		}

		fmt.Printf(element)
		p.nodes[element].PrintStructure(deep + 1)
	}
}
