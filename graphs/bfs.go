package graphs

/*
* put the root to the queue
* checking queue is empty or not
* if not -> Pop the first item -> add item's childreen to the queue
* repeat step 2 until the queue is empty
 */

// queue = []*node{}
// queue = append(queue, n)
// queue[0] // A
// queue = queue[1:]
// queue = append(queue, Node.Children...)
//

func (n *Node) BreadthFirstSearch(array []string) []string {
	// Write your code here.
	queue := make([]*Node, 0)
	if n == nil {
		return array
	}
	queue = append(queue, n)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		array = append(array, currentNode.Name)
		if currentNode.Children != nil {
			queue = append(queue, currentNode.Children...)
		}
	}
	return array
}
