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

	ids := make([]int, 0, 10)
	recordNodes := make(map[int]*Node, 10)
	for _, record := range records {
		// check for errors
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New(errorRootNodeHasParent)
		}

		if record.ID != 0 && record.ID <= record.Parent {
			return nil, errors.New(errorRecordIDLessThanParent)
		}

		if _, ok := recordNodes[record.ID]; ok {
			return nil, errors.New(errorRecordIDNotUnique)
		}

		// record is valid, create a node
		ids = append(ids, record.ID)
		recordNodes[record.ID] = &Node{ID: record.ID}
	}

	// check for a root note
	if _, ok := recordNodes[0]; !ok {
		return nil, errors.New(errorRootNodeNotFound)
	}

	// check for continuity
	sort.Ints(ids)
	for index := 0; index < len(ids); index++ {
		if ids[index] != index {
			return nil, errors.New(errorRecordsIDsMustBeContinuous)
		}
	}

	// connect the children nodes to the parents
	for _, record := range records {
		recordNode := recordNodes[record.ID]
		parentNode := recordNodes[record.Parent]

		if recordNode.ID == 0 {
			continue
		}

		parentNode.Children = append(parentNode.Children, recordNode)

		sort.Slice(parentNode.Children, func(i, j int) bool {
			return parentNode.Children[i].ID < parentNode.Children[j].ID
		})
	}

	// check for cycle
	if containsCycle(recordNodes[0]) {
		return nil, errors.New(errorTreeHasCycle)
	}

	return recordNodes[0], nil
}

func containsCycle(root *Node) bool {
	queue := append(make([]*Node, 0, 10), root)
	visited := make(map[int]bool, 10)
	for {
		focus := queue[0]
		queue = queue[1:]

		if visited[focus.ID] {
			return true
		}

		visited[focus.ID] = true
		queue = append(queue, focus.Children...)

		if len(queue) == 0 {
			break
		}
	}

	return false
}
