package main

import (
	"fmt"
	"sort"
)

func main() {
	inputArr := []int{-1,0,1,2,-1,4}
	targetSum := 0

	fmt.Println("result is: ", threeSumOptimized(inputArr, targetSum))
	
}

// Optimized Approach
func threeSumOptimized(nums []int, sum int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {

		left := i + 1 
		right := len(nums)-1

		target := sum - nums[i]
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i] ,nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1]{
					left++
				}
				for left < right && nums[right] == nums[right-1]{
					right --
				}
				left++
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
}

	return result
}