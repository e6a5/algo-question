package linkedlist

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	if linkedList == nil {
		return nil
	}
	head := linkedList
	currentPoint := head
	nextPoint := head.Next
	for nextPoint != nil {
		if currentPoint.Value == nextPoint.Value {
			nextPoint = nextPoint.Next
			continue
		}
		if currentPoint.Value != nextPoint.Value {
			currentPoint.Next = nextPoint
			currentPoint = nextPoint
			nextPoint = currentPoint.Next
		}
	}
	currentPoint.Next = nil
	return head
}
