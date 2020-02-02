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
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := make(map[int]*Node, len(records))
	for i, record := range records {
		rootError := record.ID == 0 && record.Parent != 0
		idError := record.ID != 0 && record.ID <= record.Parent
		continuityError := record.ID != i
		if rootError || idError || continuityError {
			return nil, errors.New("record error")
		}
		nodes[i] = &Node{ID: i}
		if record.ID != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
