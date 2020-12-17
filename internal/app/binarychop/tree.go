package binarychop

type BinaryNode struct {
	value int
	index int
	left  *BinaryNode
	right *BinaryNode
}

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

type BinaryTree struct {
	tree *BinaryNode
}

func (tree *BinaryTree) BinaryChop(value int, tab []int) int {
	tree.tree = buildTree(tab, 0, len(tab)-1)
	resultNode := tree.findValue(value)
	if resultNode == nil {
		return -1
	}
	return resultNode.index
}

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
