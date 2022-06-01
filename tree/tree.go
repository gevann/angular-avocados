package tree

import (
	"fmt"
)

type Tree struct {
	rightChild *Tree
	leftChild  *Tree
	parent     *Tree

	height    int
	leafCount int
	value     string
}

func (tree *Tree) GetLeafCount() int {
	var zeroValue int

	if tree == nil {
		return 0
	}

	if tree.leafCount == zeroValue {
		return 0
	}

	return tree.leafCount
}

func (tree *Tree) Insert(value string) *Tree {
	return insert(value, 0, tree, tree)
}

func (tree *Tree) NextByLeafCount() *Tree {
	var next *Tree

	if tree.rightChild != nil {
		next = tree.rightChild
	}

	for _, child := range []*Tree{tree.leftChild, tree.rightChild} {
		if child != nil && child.leafCount > next.leafCount {
			next = child
		}
	}

	return next
}

func (tree *Tree) nextByMinLeafCount() *Tree {
	var next *Tree

	if tree.leftChild == nil {
		return tree.rightChild
	} else {
		next = tree.leftChild
	}

	if tree.rightChild != nil && tree.rightChild.GetLeafCount() < next.GetLeafCount() {
		next = tree.rightChild
	}

	return next
}

func (tree *Tree) MaxByLeafCount() *Tree {
	curr := tree
	next := tree

	for next != nil {
		curr = next
		next = next.NextByLeafCount()
		curr.NextByLeafCount()
	}
	return curr
}

func (tree *Tree) MinByLeafCount() *Tree {
	curr := tree
	next := tree

	for next != nil {
		curr = next
		next = curr.nextByMinLeafCount()
	}
	return curr
}

func (tree *Tree) BreadthFirst() []string {
	var result []string

	queue := []*Tree{tree}
	var emptyValue string

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.value != emptyValue {
			result = append(result, node.value)
		}

		for _, child := range []*Tree{node.leftChild, node.rightChild} {
			if child != nil {
				queue = append(queue, child)
			}
		}
	}

	return result
}

func (tree *Tree) ToString() string {
	value := tree.value
	if value == "" {
		value = "nil"
	}
	return fmt.Sprintf("[data: %s | leaves: %d | height: %d]", value, tree.leafCount, tree.height)
}

func (tree *Tree) Print() string {
	if tree == nil {
		return "nil"
	}

	maxHeight := -1
	queue := []*Tree{tree}

	result := ""

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			fmt.Println("nil")
			continue
		}

		if node.height > maxHeight {
			maxHeight = node.height
			result += "\n"
		}

		result += fmt.Sprintf("%s", node.ToString())

		for _, child := range []*Tree{node.leftChild, node.rightChild} {
			if child != nil {
				queue = append(queue, child)
			}
		}
	}

	return result
}

func insert(value string, index int, tree *Tree, root *Tree) *Tree {
	tree.leafCount++
	// Base case: if we are on the last character of the binary string, we are at the node to insert.
	if index == len(value) {
		tree.value = value
		tree.leafCount = 0
		return root
	}

	tree.height = index

	// Recursive case
	switch value[index] {
	case '0':
		if tree.leftChild == nil {
			tree.leftChild = &Tree{parent: tree, leafCount: 0}
		}
		return insert(value, index+1, tree.leftChild, root)
	case '1':
		if tree.rightChild == nil {
			tree.rightChild = &Tree{parent: tree, leafCount: 0}
		}
		return insert(value, index+1, tree.rightChild, root)
	default:
		panic("Invalid character in binary string")
	}
}
