package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	fmt.Println(quickMaths(2, 3))

	fmt.Println(swapWords("black", "yellow"))

	simpleLoop()
}

func quickMaths(a, b int) int {
	return a * b
}

func swapWords(a, b string) (string, string) {
	return b, a
}

func simpleLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += i
	}
}