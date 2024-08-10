package main

import (
	"fmt"
	"math"
)
func main() {
	arr := []int{7,1,5,3,6,4}
	fmt.Println(findMaxProfit(arr))
}

func findMaxProfit(arr []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0
	for _, val := range arr {
		if val < minPrice {
			minPrice = val
		}

		profit := val - minPrice
		if profit > maxProfit {
			maxProfit = val
		}
	}

	return maxProfit
}