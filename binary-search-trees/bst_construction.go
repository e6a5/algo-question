package binarysearchtrees

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	if tree == nil {
		tree = &BST{
			Value: value,
		}
	}
	if tree.Value > value {
		if tree.Left == nil {
			tree.Left = &BST{
				Value: value,
			}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{
				Value: value,
			}
		} else {
			tree.Right.Insert(value)
		}
	}
	// Do not edit the return statement of this method.
	return tree
}

func (tree *BST) Contains(value int) bool {
	// Write your code here.
	if tree == nil {
		return false
	}
	if tree.Value == value {
		return true
	}
	if tree.Value > value {
		if tree.Left == nil {
			return false
		}
		return tree.Left.Contains(value)
	} else {
		if tree.Right == nil {
			return false
		}
		return tree.Right.Contains(value)
	}
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	tree.removeHelper(nil, value)
	// Do not edit the return statement of this method.
	return tree
}

func (tree *BST) removeHelper(root *BST, value int) {
	// Write your code here.
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.removeHelper(tree, value)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.removeHelper(tree, value)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tempNode := minValued(tree.Right)
			tree.Value = tempNode.Value
			tree.Right.removeHelper(tree, tree.Value)
		} else if root == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Right = tree.Right.Right
				tree.Left = tree.Right.Left
			}
		} else if root.Left == tree {
			if tree.Left != nil {
				root.Left = tree.Left
			} else {
				root.Left = tree.Right
			}
		} else if root.Right == tree {
			if tree.Left != nil {
				root.Right = tree.Left
			} else {
				root.Right = tree.Right
			}
		}
	}
}

func minValued(root *BST) *BST {
	temp := root
	for temp != nil && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}
