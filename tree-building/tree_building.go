package tree

import "sort"
import "errors"

// Record represents a post in the web forum
type Record struct {
	ID     int
	Parent int
}

// Node a struct representing a tree node in the reconstructed post tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build function receives a slice of Record structs
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	inserted, nodes := make([]bool, len(records)), make([]*Node, len(records))
	for i, record := range records {
		rootError := record.ID == 0 && record.Parent != 0
		idError := record.ID != 0 && record.ID <= record.Parent
		continuityError := record.ID != i
		if rootError || idError || continuityError {
			return nil, errors.New("record error")
		}
		nodes[i], inserted[record.ID] = &Node{ID: i}, true
		if record.ID != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[i])
		}
	}
	for _, node := range nodes {
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].ID < node.Children[j].ID
		})
	}
	return nodes[0], nil
}
