package main

import "fmt"

func main() {
	inputArr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(inputArr))
	inputArr1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(anotherOptimizedSolution(inputArr1))
}

func removeDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	key := nums[0]
	count := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] == key {
			continue
		} else {
			key = nums[j]
			nums[count] = nums[j]
			count++
		}
	}

	return nums[:count]
}

func anotherOptimizedSolution(nums []int) []int {
	if len(nums) == 0{
		return []int{}
	}

	i := 0
	for j := 1; j < len(nums); j++{
		fmt.Println(nums[i], nums[j])
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return nums[:i+1]
}