package binarysearchtrees

import "math"

func (tree *BST) ValidateBst() bool {
	// Write your code here.
	return tree.validate(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validate(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.validate(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validate(tree.Value, max) {
		return false
	}
	return true
}
