package main

import "fmt"

func main() {
	arr := []int{1,2,3}
	//ans := [][]int{}

	generatePermutation(arr, 0)
	fmt.Println(ans)

}

var ans [][]int

func generatePermutation(nums []int, index int) {
	if index == len(nums) {
		ans = append(ans, nums)
		return
	}

	for i := index; i < len(nums); i++ {
		swap(nums, i, index)
		generatePermutation(nums, index+1)
		swap(nums, index, i)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}