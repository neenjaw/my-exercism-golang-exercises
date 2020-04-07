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
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("record error")
		}
		nodes[i] = &Node{ID: i}
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
