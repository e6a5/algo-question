package binarytrees

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	// Write your code here.
	helper(tree)
}

func helper(tree *BinaryTree) *BinaryTree {
	if tree == nil {
		return nil
	}
	invertedL := helper(tree.Left)
	invertedR := helper(tree.Right)
	tree.Left = invertedR
	tree.Right = invertedL
	return tree
}
