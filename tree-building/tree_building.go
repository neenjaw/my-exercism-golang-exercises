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
	// if no records, abort early
	if len(records) == 0 {
		return nil, nil
	}

	inserted, nodes := make([]bool, len(records)), make([]Node, len(records))
	for _, record := range records {
		rootError := record.ID == 0 && record.Parent != 0
		idError := record.ID != 0 && record.ID <= record.Parent
		continuityError := record.ID >= len(records)
		if rootError || idError || continuityError {
			return nil, errors.New("record error")
		}
		if inserted[record.ID] {
			return nil, errors.New("duplicate record error")
		}
		nodes[record.ID].ID, inserted[record.ID] = record.ID, true
		if record.ID != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[record.ID])
		}
	}
	for _, node := range nodes {
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].ID < node.Children[j].ID
		})
	}
	return &nodes[0], nil
}
