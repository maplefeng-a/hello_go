package main

import (
	"fmt"
)
func square(n int) int { return n * n }
func cube(n int, f func(int) int) int {
	return n * f(n)
}
func main() {
	f := cube
	fmt.Println(f(3, square))
}
