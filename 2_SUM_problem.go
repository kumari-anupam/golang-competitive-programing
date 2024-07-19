package main

import (
	"fmt"
	"sort"
)

func main() {
	inputArr := []int{-1,0,1,2,-1,4}
	targetSum := 0

	fmt.Println(len(inputArr))

	naiveOuput := twoSum(inputArr, targetSum)
	fmt.Println("Naive result is: ", naiveOuput)

	optimizedOuput := twoSumOptimized(inputArr, targetSum)
	fmt.Println("Optimized result is: ", optimizedOuput)
}

// NaiveApproach
func twoSum(nums []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{nums[i], nums[j]})
			}
		}
	}

	return result
}

// Optimized Approach
func twoSumOptimized(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	left := 0
	right := len(nums)-1

	for left < right {
		if nums[left]+nums[right] == target {
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}

	return result
}