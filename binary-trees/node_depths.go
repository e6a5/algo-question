package main

func NodeDepths(root *BinaryTree) int {
	return SumNodeDepths(root, 0)
}

func SumNodeDepths(root *BinaryTree, depth int) int {
	sum := depth
	if root.Left != nil {
		sum += SumNodeDepths(root.Left, depth+1)
	}
	if root.Right != nil {
		sum += SumNodeDepths(root.Right, depth+1)
	}
	return sum
}
