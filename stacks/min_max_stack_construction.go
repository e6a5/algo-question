package stacks

import "math"

type MinMaxStack struct {
	// Write your code here.
	data   []int
	min    int
	max    int
	length int
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	return &MinMaxStack{
		data:   make([]int, 0),
		min:    math.MaxInt,
		max:    math.MinInt,
		length: 0,
	}
}

func (stack *MinMaxStack) Peek() int {
	// Write your code here.
	if stack.length > 0 {
		return stack.data[stack.length-1]
	}
	return -1
}

func (stack *MinMaxStack) Pop() int {
	// Write your code here.
	if stack.length > 0 {
		newLength := stack.length - 1
		topNum := stack.data[newLength]
		stack.data = stack.data[:newLength]
		stack.length = newLength
		if topNum == stack.min {
			stack.min = findMin(stack.data, stack.length)
		}
		if topNum == stack.max {
			stack.max = findMax(stack.data, stack.length)
		}
		return topNum
	}
	return -1
}

func (stack *MinMaxStack) Push(number int) {
	// Write your code here.
	if number < stack.min {
		stack.min = number
	}
	if number > stack.max {
		stack.max = number
	}
	stack.data = append(stack.data, number)
	stack.length = stack.length + 1
}

func (stack *MinMaxStack) GetMin() int {
	// Write your code here.
	return stack.min
}

func (stack *MinMaxStack) GetMax() int {
	// Write your code here.
	return stack.max
}

func findMin(data []int, n int) int {
	if n == 0 {
		return math.MaxInt
	}
	min := data[0]
	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
	}
	return min
}

func findMax(data []int, n int) int {
	if n == 0 {
		return math.MinInt
	}
	max := data[0]
	for i := 1; i < n; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}
