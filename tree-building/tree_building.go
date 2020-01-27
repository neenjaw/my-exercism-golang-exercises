package tree

import "sort"
import "errors"

const (
	errorRootNodeNotFound           string = "no root node"
	errorRootNodeHasParent          string = "root cannot have parent"
	errorRecordIDLessThanParent     string = "record id must be greater than parent id"
	errorRecordIDNotUnique          string = "record id's must be unique"
	errorRecordsIDsMustBeContinuous string = "ids must be zero-indexed and continuous"
	errorTreeHasCycle               string = "id cycle detected"
)

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

	inserted := make([]bool, len(records))
	nodes := make([]Node, len(records))
	for _, record := range records {
		// check for errors
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New(errorRootNodeHasParent)
		}
		if record.ID != 0 && record.ID <= record.Parent {
			return nil, errors.New(errorRecordIDLessThanParent)
		}
		if record.ID >= len(records) {
			return nil, errors.New(errorRecordsIDsMustBeContinuous)
		}
		if inserted[record.ID] {
			return nil, errors.New(errorRecordIDNotUnique)
		}

		// record is valid, create a node
		nodes[record.ID].ID = record.ID
		inserted[record.ID] = true

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
