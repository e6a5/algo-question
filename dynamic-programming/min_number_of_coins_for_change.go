package dynamicprogramming

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Write your code here.
	changes := make([]int, n+1)
	for i := range changes {
		changes[i] = math.MaxInt32
	}
	changes[0] = 0
	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount++ {
			if denom <= amount {
				changes[amount] = min(changes[amount], 1+changes[amount-denom])
			}
		}
	}
	// if changes at n equals to math max int32
	// it means we can not make change for that target amount using the provided coins
	if changes[n] != math.MaxInt32 {
		return changes[n]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
