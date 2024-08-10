package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abcabcdbb"
	fmt.Println(longestSubsequence(s))
}

func longestSubsequence(s string) int {
	// generate all subsequence
	// then find the longest one
	//take map when traversing on subsequence store it in map if map size is equal to (j-i+1) return j-i+1
	// else return max

	// complexity o(n2)

	//using sliding window

	i := 0
	j := 0
	charCount := make(map[byte]int)
	maxLength := math.MinInt

	for j < len(s) {
		if val, ok := charCount[s[j]]; ok {
			charCount[s[j]] = val + 1
		} else {
			charCount[s[j]] = 1
		}

		if len(charCount) == (j - i + 1) {
			maxLength = max(maxLength, j-i+1)
			j++
		} else if len(charCount) < (j-i+1) {
			for len(charCount) < (j-i+1) {
				if val, ok := charCount[s[i]]; ok {
					charCount[s[i]] = val - 1
					if val-1 == 0 {
						delete(charCount, s[i])
					}
				}
				i++
			}
			j++
		}
	}

	return maxLength
}