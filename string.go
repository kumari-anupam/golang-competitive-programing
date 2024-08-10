package main

import "fmt"

func main() {
	str := "Hello World"
	fmt.Println(str)
	str = str + "!"
	fmt.Println(str)

	s := stringConcate(str)
	fmt.Println(string(s))

	runes := []rune(s)
	for _, val := range runes {
		fmt.Println(string(val))
	}

	fmt.Println(string(runes[:5]))
}

func stringConcate(str string) string {
	s := ""
	for _, val := range str {
		s += string(val)
	}

	return s
}