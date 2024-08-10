package main

import "fmt"

func main() {
	output := generateParenthesis(1)
	fmt.Println(output)

	output = generateParenthesis(2)
	fmt.Println(output)

	output = generateParenthesis(3)
	fmt.Println(output)
}

func generateParenthesis(n int) []string {
	input := ""
	for i := 0; i < n; i++ {
		input = input + "  "
	}
	output := generateParenthesisUtil(0, n, 0, 0, []rune(input))
	return output
}

func generateParenthesisUtil(pos, n, open, close int, input []rune) []string {

	var output []string
	if pos == n*2 {
		output = append(output, string(input))
		return output
	}

	if close < open {
		input[pos] = ')'
		result := generateParenthesisUtil(pos+1, n, open, close+1, input)
		output = append(output, result...)

	}

	if open < n {
		input[pos] = '('
		result := generateParenthesisUtil(pos+1, n, open+1, close, input)
		output = append(output, result...)
	}

	return output
}