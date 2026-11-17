package main

import "fmt"

func main() {
	first, second := 16, 17
	var multiplier func(x, y int) int
	multiplier = func(x, y int) int { return x * y }
	fmt.Println(multiplier(first, second))
	divider := func(x, y int) int { return x / y }
	fmt.Println(divider(second, first)) 
}
