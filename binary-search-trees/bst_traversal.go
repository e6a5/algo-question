package binarysearchtrees

func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.
	// LeftNode -> Root -> RightNode
	if tree == nil {
		return []int{}
	}
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.
	// root -> left -> right
	if tree == nil {
		return []int{}
	}
	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	// left -> right -> root
	if tree == nil {
		return []int{}
	}
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)
	return array
}
