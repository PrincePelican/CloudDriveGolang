package service

import (
	"cloud-service/DTO"
	"cloud-service/entity"
	"testing"
)

func compareNodeNames(node1, node2 *DTO.DirectoryNode) bool {
	if node1.Name != node2.Name {
		return false
	}

	for name, child1 := range node1.Nodes {
		child2, ok := node2.Nodes[name]
		if !ok || !compareNodeNames(child1, child2) {
			return false
		}
	}

	return true
}

func compareResourceEntity(entity1, entity2 entity.ResourceEntity) bool {
	if entity1.Name != entity2.Name || entity1.Key != entity2.Key || entity1.ResourceType != entity2.ResourceType {
		return false
	}

	if len(entity1.Childs) != len(entity2.Childs) {
		return false
	}

	for i := range entity1.Childs {
		if !compareResourceEntity(entity1.Childs[i], entity2.Childs[i]) {
			return false
		}
	}

	return true
}

func TestEmptyPathsToTreeStructure(t *testing.T) {
	expected := &DTO.DirectoryNode{
		Name:  "node1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	paths := []string{}
	actual, _ := ConvertFromPathsToTreeStructure(paths, "node1")
	if !compareNodeNames(actual, expected) {
		t.Errorf("Structure should be the same")
	}
}

func TestFromDeepPathsToTreeStructure(t *testing.T) {
	node1 := &DTO.DirectoryNode{
		Name:  "node1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}

	node2 := &DTO.DirectoryNode{
		Name:  "node2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2.Nodes["node1"] = node1

	node3 := &DTO.DirectoryNode{
		Name:  "node3",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node3.Nodes["node2"] = node2

	expected := &DTO.DirectoryNode{
		Name:  "node4",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	expected.Nodes["node3"] = node3
	paths := []string{"node3/node2/node1"}
	actual, _ := ConvertFromPathsToTreeStructure(paths, "node4")
	if !compareNodeNames(actual, expected) {
		t.Errorf("Structure should be the same")
	}
}

func TestWidthPathsToTreeStructure(t *testing.T) {
	node1_1 := &DTO.DirectoryNode{
		Name:  "node1_1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node1_2 := &DTO.DirectoryNode{
		Name:  "node1_2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node1_3 := &DTO.DirectoryNode{
		Name:  "node1_3",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2 := &DTO.DirectoryNode{
		Name:  "node2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2.Nodes["node1_1"] = node1_1
	node2.Nodes["node1_2"] = node1_2
	node2.Nodes["node1_3"] = node1_3

	expected := &DTO.DirectoryNode{
		Name:  "node3",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	expected.Nodes["node2"] = node2
	paths := []string{"node2/node1_1", "node2/node1_2", "node2/node1_3"}
	actual, _ := ConvertFromPathsToTreeStructure(paths, "node3")
	if !compareNodeNames(actual, expected) {
		t.Errorf("Structure should be the same")
	}
}

func TestCheckIfLastNodeHasKeyInTreeStructure(t *testing.T) {
	paths := []string{"nodeFile"}
	actual, keys := ConvertFromPathsToTreeStructure(paths, "node1")
	if actual.Nodes["nodeFile"].Key == "" {
		t.Errorf("Node should have key")
	}
	if keys[0] != actual.Nodes["nodeFile"].Key {
		t.Errorf("ConvertFromPathsToTreeStructure should return the same key that is in the node")
	}
}

func TestConvertingSingleDirNodeToResourceStructure(t *testing.T) {
	node1 := &DTO.DirectoryNode{
		Name:  "node1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	expectedEntityNode1 := &entity.ResourceEntity{
		Name:   "node1",
		Childs: make([]entity.ResourceEntity, 0),
	}

	actual := ConvertFromDirStructureToResourceTree(node1)

	if !compareResourceEntity(*actual, *expectedEntityNode1) {
		t.Errorf("Structure should be the same")
	}
}

func TestTheSameDirNodeToResourceStructure(t *testing.T) {
	node1 := &DTO.DirectoryNode{
		Name:         "node1",
		Key:          "KLUCZ",
		ResourceType: 0,
		Nodes:        make(map[string]*DTO.DirectoryNode),
	}
	expectedEntityNode1 := &entity.ResourceEntity{
		Name:         "node1",
		Key:          "KLUCZ",
		ResourceType: 0,
		Childs:       make([]entity.ResourceEntity, 0),
	}

	actual := ConvertFromDirStructureToResourceTree(node1)

	if !compareResourceEntity(*actual, *expectedEntityNode1) {
		t.Errorf("Structure should be the same")
	}
}

func TestConvertingDeepDirNodeToResourceStructure(t *testing.T) {
	node1 := &DTO.DirectoryNode{
		Name:  "node1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}

	node2 := &DTO.DirectoryNode{
		Name:  "node2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2.Nodes["node1"] = node1

	node3 := &DTO.DirectoryNode{
		Name:  "node3",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node3.Nodes["node2"] = node2

	node4 := &DTO.DirectoryNode{
		Name:  "node4",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node4.Nodes["node3"] = node3

	expectedEntityNode1 := &entity.ResourceEntity{
		Name:   "node1",
		Childs: make([]entity.ResourceEntity, 0),
	}
	expectedEntityNode2 := &entity.ResourceEntity{
		Name:   "node2",
		Childs: []entity.ResourceEntity{*expectedEntityNode1},
	}
	expectedEntityNode3 := &entity.ResourceEntity{
		Name:   "node3",
		Childs: []entity.ResourceEntity{*expectedEntityNode2},
	}
	expectedEntityNode4 := &entity.ResourceEntity{
		Name:   "node4",
		Childs: []entity.ResourceEntity{*expectedEntityNode3},
	}

	actual := ConvertFromDirStructureToResourceTree(node4)

	if !compareResourceEntity(*actual, *expectedEntityNode4) {
		t.Errorf("Structure should be the same")
	}
}

func TestConvertingWidthDirNodeToResourceStructure(t *testing.T) {
	node1_1 := &DTO.DirectoryNode{
		Name:  "node1_1",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node1_2 := &DTO.DirectoryNode{
		Name:  "node1_2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node1_3 := &DTO.DirectoryNode{
		Name:  "node1_3",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2 := &DTO.DirectoryNode{
		Name:  "node2",
		Nodes: make(map[string]*DTO.DirectoryNode),
	}
	node2.Nodes["node1_1"] = node1_1
	node2.Nodes["node1_2"] = node1_2
	node2.Nodes["node1_3"] = node1_3

	expectedEntityNode1_1 := &entity.ResourceEntity{
		Name:   "node1_1",
		Childs: make([]entity.ResourceEntity, 0),
	}
	expectedEntityNode1_2 := &entity.ResourceEntity{
		Name:   "node1_2",
		Childs: make([]entity.ResourceEntity, 0),
	}
	expectedEntityNode1_3 := &entity.ResourceEntity{
		Name:   "node1_3",
		Childs: make([]entity.ResourceEntity, 0),
	}
	expectedEntityNode2 := &entity.ResourceEntity{
		Name:   "node2",
		Childs: []entity.ResourceEntity{*expectedEntityNode1_1, *expectedEntityNode1_2, *expectedEntityNode1_3},
	}

	actual := ConvertFromDirStructureToResourceTree(node2)

	if !compareResourceEntity(*actual, *expectedEntityNode2) {
		t.Errorf("Structure should be the same")
	}
}
