package main

import "fmt"

type stack struct {
	item []int
}

func (s *stack) push(x int) {
	s.item = append(s.item, x)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return 0
	}

	x := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]

	return x
}

func (s *stack) top() int {
	x := s.item[len(s.item)-1]

	return x
}

func (s *stack) print() {
	for i := 0; i < len(s.item); i++ {
		fmt.Print(s.item[i], " ")
	}
}

func (s *stack) isEmpty() bool {
	return len(s.item) == 0
}

func main() {
	s := stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.pop()
	s.push(6)
	fmt.Println(s.top())
	s.print()
}