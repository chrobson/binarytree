package main

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// 1. calculate number of nodes without childrens

func countLeafNodes(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return countLeafNodes(node.Left) + countLeafNodes(node.Right)
}

// 2. calulate number of edges
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPath(node *Node) int {
	if node == nil {
		return 0
	}
	leftPath := longestPath(node.Left)
	rightPath := longestPath(node.Right)
	return 1 + max(leftPath, rightPath)
}

// 3. check if the tree instances are equivalent
func areTreesEquivalent(tree1, tree2 *Node) bool {
	// Both trees are empty, so they are equivalent.
	if tree1 == nil && tree2 == nil {
		return true
	}

	// One of the trees is empty, so they are not equivalent.
	if tree1 == nil || tree2 == nil {
		return false
	}

	// Check if the current nodes have the same value and if their children are equivalent.
	return tree1.Value == tree2.Value &&
		areTreesEquivalent(tree1.Left, tree2.Left) &&
		areTreesEquivalent(tree1.Right, tree2.Right)
}

func main() {
	// Create the nodes
	root := &Node{Value: 5}
	leftChild := &Node{Value: 3}
	leftLeftChild := &Node{Value: 2}
	leftRightChild := &Node{Value: 5}
	rightChild := &Node{Value: 7}
	rightLeftChild := &Node{Value: 1}
	rightRightChild := &Node{Value: 0}
	rightRightLeftChild := &Node{Value: 2}
	rightRightRightChild := &Node{Value: 8}
	rightRightRightLeftChild := &Node{Value: 5}

	// Connect the nodes according to the provided structure.
	root.Left = leftChild
	root.Right = rightChild

	leftChild.Left = leftLeftChild
	leftChild.Right = leftRightChild

	rightChild.Left = rightLeftChild
	rightChild.Right = rightRightChild

	rightRightChild.Left = rightRightLeftChild
	rightRightChild.Right = rightRightRightChild

	rightRightRightChild.Left = rightRightRightLeftChild

	// Print the tree.
	//printTree(root, 0)
	fmt.Println(countLeafNodes(root))
	// exclude root to show the edges instead of nodes
	longestPathEdges := longestPath(root) - 1
	fmt.Println(longestPathEdges)
	fmt.Println(areTreesEquivalent(root, rightLeftChild))
	printTree(rightChild, 0)
	// create same tree as nested struct
	root2 := &Node{5,
		&Node{3,
			&Node{2, nil, nil},
			&Node{5, nil, nil}},
		&Node{7,
			&Node{1, nil, nil},
			&Node{0,
				&Node{2, nil, nil},
				&Node{8,
					&Node{5, nil, nil},
					nil}}}}
	//printTree(root2, 0)
	fmt.Println(areTreesEquivalent(root, root2))
}

// helper allowing to print the tree and visulize it in the console (working with whole tree and given instance)
func printTree(node *Node, level int) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "    "
	}
	format += "---[ "

	level++
	printTree(node.Left, level)
	fmt.Printf(format+"%d\n", node.Value)
	printTree(node.Right, level)
}
