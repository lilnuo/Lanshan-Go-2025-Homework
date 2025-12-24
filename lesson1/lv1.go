package main

import "fmt"

const pi = 3.14

func main() {
	r := 5

	area := pi * float64(r) * float64(r)
	fmt.Println("面积:", area)
}
