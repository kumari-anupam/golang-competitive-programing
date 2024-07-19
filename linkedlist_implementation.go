package main

import "fmt"

type node struct {
	data int
	next *node
}

type singleLinkedList struct {
	head *node
}

func (s *singleLinkedList) insertAtEnd(data int) {
	if s.head == nil {
		s.head = &node{data: data}
		return
	}

	current := s.head
	for current.next != nil {
		current = current.next
	}

	current.next = &node{data: data}
}

func (s *singleLinkedList) insertAtBeginning(data int) {
	newNode := &node{data: data, next: s.head}
	s.head = newNode
}

func (s *singleLinkedList) insertAfterGivenNodeValue(nodeValue int, data int) {
	current := s.head

	if s.head == nil {
		fmt.Printf("%d not found\n", nodeValue)
		return
	}

	for current != nil && current.data != nodeValue {
		current = current.next
	}

	if current.data == nodeValue {
		newNode := &node{data: data, next: current.next}
		current.next = newNode
		return
	}

	fmt.Printf("%d not found\n", nodeValue)

}

func (s *singleLinkedList) insertBeforeGivenNodeValue(nodeValue int, data int) {
	prev := &node{}
	curr := s.head

	if s.head == nil {
		fmt.Printf("%d not found\n", nodeValue)
		return
	}

	if curr.data == nodeValue {
		newNode := &node{data: data, next: curr}
		s.head = newNode
		return
	}

	for curr != nil && curr.data != nodeValue {
		prev = curr
		curr = curr.next
	}

	if curr.data == nodeValue {
		newNode := &node{data: data, next: curr}
		prev.next = newNode
		return
	}

	fmt.Printf("%d not found\n", nodeValue)
}

func (s *singleLinkedList) deleteAtBeginning() {
	if s.head == nil {
		return
	}

	if s.head.next == nil {
		s.head = nil
		return
	}

	if s.head.next != nil {
		s.head = s.head.next
		return
	}

}

func (s *singleLinkedList) deleteAtEnd() {
	curr := s.head
	prev := &node{}

	if s.head.next == nil {
		s.head  = nil
		return
	}

	for curr.next != nil {
		prev = curr
		curr = curr.next
	} 

	prev.next = nil
}

func (s *singleLinkedList) deleteBeforeGivenNodeVal(nodeVal int) {
	if s.head == nil {
		return
	}

	if s.head.data == nodeVal {
		s.head = s.head.next
		return
	}

	curr := s.head
	prev := &node{}

	for curr != nil && curr.next != nil && curr.next.data != nodeVal {
		prev = curr
		curr = curr.next
	}

	if curr.next != nil && curr.next.data == nodeVal {
		prev.next = curr.next
		return
	}

	fmt.Printf("%d not found\n", nodeVal)
}

func (s *singleLinkedList) deleteAfterGivenNodeVal(nodeVal int) {
	if s.head.next == nil &&s.head.data == nodeVal {
		s.head = nil
		return
	}

	curr := s.head

	for curr != nil && curr.next != nil && curr.next.data != nodeVal {
		curr = curr.next
	}

	if curr.next != nil && curr.next.data == nodeVal {
		curr.next = curr.next.next
		return
	}

	fmt.Printf("%d not found\n", nodeVal)
}



func (s *singleLinkedList) print() {
	curr := s.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
}

func main() {
	s := &singleLinkedList{}
	s.insertAtEnd(10)
	s.insertAtEnd(20)
	s.insertAtEnd(30)
	s.insertAtBeginning(5)
	s.insertBeforeGivenNodeValue(5, 2)
	s.insertAfterGivenNodeValue(10, 15)
	s.insertBeforeGivenNodeValue(30, 25)
	s.print()

	fmt.Println("\ndeleting at beginning")
	s.deleteAtBeginning()
	fmt.Println("deleting at end")
	s.deleteAtEnd()
	fmt.Println("deleting before given value 20")
	s.deleteBeforeGivenNodeVal(20)
	fmt.Println("deleting after given value 15")
	s.deleteAfterGivenNodeVal(15)
	s.print()
}