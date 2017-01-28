package main

import (
	"fmt"
)

func (n *Rect) String() string {
	rect := "Rect [ " + fmt.Sprint(n.x1) + " , " + fmt.Sprint(n.y1) + " , " + fmt.Sprint(n.x2) + " , " + fmt.Sprint(n.y2) + " ] area = " + fmt.Sprint(n.area)
	return rect
	//return fmt.Sprint(n.area)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Rect
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Rect) {

	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Rect {

	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Rect, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Rect
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Rect) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Rect, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Rect {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
func (q *Queue) IsEmpty() bool {
	if q.count <= 0 {
		return true
	} else {
		return false
	}
}
