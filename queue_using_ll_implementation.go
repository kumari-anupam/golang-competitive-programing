package main

import "fmt"

type qnode struct {
	data int
	next *qnode
}

type Queue struct {
	front *qnode
	rear  *qnode
}

func (q *Queue) enqueue_qnode(element int) {
	// queue is empty
	newNode := &qnode{data: element}
	if q.rear == nil {
		newNode := &qnode{data: element, next: nil}
		q.rear, q.front = newNode, newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) dequeue_qnode() int {
	if q.front == nil {
		return 0
	}

	temp := q.front
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return temp.data
}

func (q *Queue) print() {
	curr := q.front
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
}

func main() {
	q := &Queue{}
	q.enqueue_qnode(1)
	q.enqueue_qnode(2)
	q.dequeue_qnode()
	q.dequeue_qnode()
	q.enqueue_qnode(3)
	q.enqueue_qnode(4)
	q.enqueue_qnode(5)
	q.enqueue_qnode(6)
	q.print()
}