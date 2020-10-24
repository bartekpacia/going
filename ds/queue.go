package ds

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(value int) {
	q.nodes[q.tail] = &Node{value: value}
	q.tail = (q.tail + 1)
	q.count++
}
