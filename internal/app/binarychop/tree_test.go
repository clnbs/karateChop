package binarychop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree_BinaryChop(t *testing.T) {
	GeneralTestInterfaceImplementation(&BinaryTree{}, t)
}

func TestBinaryTree_BinaryChop2(t *testing.T) {
	ShuffleGeneralTestInterface(&BinaryTree{}, t)
}

func TestBinaryTree_BinaryChop3(t *testing.T) {
	GeneralTestInterface(&BinaryTree{}, t)
}

func TestBinaryNode_insert(t *testing.T) {
	head := &BinaryNode{
		value: 5,
		index: 0,
		left:  nil,
		right: nil,
	}
	head.insert(10, 1)
	assert.Equal(t, 10, head.right.value)

	head.insert(3, 2)
	assert.Equal(t, 3, head.left.value)

	head.insert(8, 3)
	assert.Equal(t, 8, head.right.left.value)

	head.insert(1, 4)
	assert.Equal(t, 1, head.left.left.value)

	head = nil
	head.insert(5, 0)
	assert.Nil(t, head)
}

func TestBinaryTree_buildTree(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	tree := &BinaryTree{}
	tree.tree = buildTree(values, 0, len(values)-1)

	assert.Equal(t, 4, tree.tree.value)
	//left part of the tree
	assert.Equal(t, 2, tree.tree.left.value)
	assert.Equal(t, 1, tree.tree.left.left.value)
	assert.Equal(t, 3, tree.tree.left.right.value)
	//right part of the tree
	assert.Equal(t, 6, tree.tree.right.value)
	assert.Equal(t, 7, tree.tree.right.right.value)
	assert.Equal(t, 5, tree.tree.right.left.value)
}
