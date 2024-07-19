package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("Initial: len=%d cap=%d %v\n", len(s), cap(s), s)

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("After appending %d: len=%d cap=%d %v\n", i, len(s), cap(s), s)
	}

	data := []int{2, 3, 4, 5, 6, 7, 8, 9, 20}
	fmt.Printf("len=%d cap=%d %v\n", len(data), cap(data), data)

	// Avoid doing this in a tight loop
	for i := 0; i < len(data); i++ {
		subSlice := data[i : i+2]
		// process subSlice
		fmt.Printf("len=%d cap=%d %v\n", len(subSlice), cap(subSlice), subSlice, &subSlice)
	}
}