package binarychop

// BinaryNode represent a value as a node in a binary tree
type BinaryNode struct {
	value int
	index int
	left  *BinaryNode
	right *BinaryNode
}

// insert can be use to insert a new value in a tree already built.
// At each node, it compares the wannabe inserted value to the node value
// and navigate in the tree to find the right spot to register the value
func (node *BinaryNode) insert(value, index int) {
	if node == nil {
		return
	}
	if value <= node.value {
		if node.left == nil {
			node.left = &BinaryNode{
				value: value,
				index: index,
				left:  nil,
				right: nil,
			}
		} else {
			node.left.insert(value, index)
		}
		return
	}

	if node.right == nil {
		node.right = &BinaryNode{
			value: value,
			index: index,
			left:  nil,
			right: nil,
		}
	} else {
		node.right.insert(value, index)
	}
}

// BinaryTree is just a wrapper of BinaryNode in order to get the binary tree
type BinaryTree struct {
	tree *BinaryNode
}

// BinaryChop implement binarychop.Algorithm interface for the binary tree algorithm
func (tree *BinaryTree) BinaryChop(value int, tab []int) int {
	tree.tree = buildTree(tab, 0, len(tab)-1)
	resultNode := tree.findValue(value)
	if resultNode == nil {
		return -1
	}
	return resultNode.index
}

// buildTree build a balance tree from a int array
func buildTree(tab []int, start, end int) *BinaryNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &BinaryNode{
		value: tab[mid],
		index: mid,
		left:  nil,
		right: nil,
	}
	root.left = buildTree(tab, start, mid-1)
	root.right = buildTree(tab, mid+1, end)
	return root
}

// findValue return the first node containing the actual value passed in params
func (tree *BinaryTree) findValue(value int) *BinaryNode {
	tmpNode := tree.tree
	for tmpNode != nil {
		if value == tmpNode.value {
			return tmpNode
		}
		if value > tmpNode.value {
			tmpNode = tmpNode.right
			continue
		}
		if value < tmpNode.value {
			tmpNode = tmpNode.left
		}
	}
	return nil
}
