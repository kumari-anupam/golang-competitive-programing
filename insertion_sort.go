package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i-1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}