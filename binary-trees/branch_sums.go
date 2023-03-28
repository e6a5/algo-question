package binarytrees

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Value}
	}
	sums := make([]int, 0)
	for _, sum := range BranchSums(root.Left) {
		sums = append(sums, sum+root.Value)
	}
	for _, sum := range BranchSums(root.Right) {
		sums = append(sums, sum+root.Value)
	}
	return sums
}
