package main

import "fmt"

func main() {
	fmt.Println(countSubstring("aabbcc", 2))
}

func countSubstring(s string, k int) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]int)
		for j := i; j < len(s); j++ {
			m[s[j]]++
			if isValid(m, k) {
				ans++
			}
		}
	}

	return ans
}

func isValid(freq map[byte]int, k int) bool {
	for _, v := range freq {
		if  v != k {
			return false
		}
	}

	return true
}

func optimizedCountSubstring(s string, k int) int {
}