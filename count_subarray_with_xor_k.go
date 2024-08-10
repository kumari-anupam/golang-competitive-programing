package main

import "fmt"

func main() {
	arr := []int{ 4, 2, 2, 6, 4 }
	k := 6
	fmt.Println(countSubarraywithGivenXOR(arr, k))
}

// optimal approach
func countSubarraywithGivenXOR(arr []int, k int) int {
	count := 0
	xor := 0
	xorMap := make(map[int]int)

	for _, val := range arr {
		xor = xor ^ val
		if xor == k {
			count ++
		}

		if val, ok := xorMap[xor ^ k]; ok {
			count += val
		}

		if val, ok := xorMap[xor]; ok {
			xorMap[xor] = val + 1
		} else {
			xorMap[xor] = 1
		}
	}

	return count
}