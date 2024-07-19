package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(mergeSort(arr))
}

func mergeSort(nums []int) []int{
	if  len(nums) <= 1{
		return nums
	}
	
	left := 0
	right := len(nums)-1
	mid := left + (right-left)/2
	leftArr := mergeSort(nums[:mid])
	rightArr := mergeSort(nums[mid+1:])

	return merge(leftArr, rightArr)
}

func merge(a1, a2 []int) []int{
	i := 0
	j := 0
	arr3 := []int{}
	k := 0
	if i < len(a1) && j < len(a2) {
		if a1[i] < a2[j]{
			arr3[k] = a1[i]
			i++
		}else {
			arr3[k] = a2[j]
			j++
		}
		k++
	}

	for i < len(a1) {
		arr3[k] = a1[i]
		k++
		i++
	}

	for j < len(a2){
		arr3[k] = a2[j]
		k++
		j++
	}

	return arr3
}