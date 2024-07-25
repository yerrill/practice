package main

import "fmt"

func main() {
	charsInSubstring := make(map[byte]int)
	charsInSubstring['a'] = 1

	fmt.Printf("a: %v, b: %v", charsInSubstring['a'], charsInSubstring['b'])
}
