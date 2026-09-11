package main

import "fmt"

func main() {
	fmt.Println("hello, ", add(2, 3))
}
func add(a int, b int) int {
	return a + b
}
