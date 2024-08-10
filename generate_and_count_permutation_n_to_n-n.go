package main

import (
	"fmt"
)

// count all permutation of n to n-n, combination of R and L
// don't include if 2 Ls are adjacent

// Helper function to calculate binomial coefficient (n choose k)
func binomial(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	k = min(k, n-k) // Take advantage of symmetry
	c := 1
	for i := 0; i < k; i++ {
		c = c * (n - i) / (i + 1)
	}
	return c
}

// Min function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function to calculate the total number of valid permutations
func countValidPermutations(n int) int {
	totalCount := 0

	// Iterate over sequence lengths from n down to 0
	for k := n; k >= 0; k-- {
		// Sum up all valid permutations for each m
		for m := 0; m <= min(k, k/2); m++ {
			totalCount += binomial(k-m+1, m)
		}
	}

	return totalCount
}

func main() {
	n := 3 // Example value, can be changed to any positive integer
	result := countValidPermutations(n)
	fmt.Printf("Total valid permutations where no two Ls are adjacent for sequences of length %d down to 0: %d\n", n, result)
}

/*
For a small example where n = 3:

k = 3: Valid sequences are RRR, RRL, RLR, LRR, so 4 valid sequences.
k = 2: Valid sequences are RR, RL, LR, so 3 valid sequences.
k = 1: Valid sequences are R, L, so 2 valid sequences.
k = 0: Valid sequence is the empty sequence.
Adding these, we get 4 + 3 + 2 + 1 = 10 valid sequences in total.
*/