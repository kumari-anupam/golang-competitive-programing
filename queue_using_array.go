package main

import "fmt"

type queue struct {
	element []int
	size    int
}

func (q *queue) enqueue(element int) {
	q.element = append(q.element, element)

	q.size++
}

func (q *queue) dequeue() int {
	if q.isEmpty() {
		return 0
	}

	element := q.element[0]
	q.element = q.element[1:]
	q.size--
	return element
}

func (q *queue) isEmpty() bool {
	return len(q.element) == 0
}

func (q *queue) print() {
	for i := 0; i < q.size; i++ {
		fmt.Print(q.element[i], " ")
	}
}

func main() {

	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	x := q.dequeue()
	fmt.Println(x)
	fmt.Println(q.dequeue())

	fmt.Println(q.isEmpty())
	q.print()
}