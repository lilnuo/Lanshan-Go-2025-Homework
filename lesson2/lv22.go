package main

import "fmt"

type operate func(float64, float64) float64

func add(a, b float64) float64 {
	return a + b
}
func sub(a, b float64) float64 {
	return a - b
}
func mul(a, b float64) float64 {
	return a * b
}
func div(a, b float64) float64 {
	return a / b
}

func calculate(x float64, y float64, op operate) float64 {
	return op(x, y)
}
func main() {
	var x, y float64
	fmt.Scanln(&x, &y)
	result := calculate(x, y, add)
	fmt.Println(result)
	result = calculate(x, y, sub)
	fmt.Println(result)
	result = calculate(x, y, mul)
	fmt.Println(result)
	result = calculate(x, y, div)
	fmt.Println(result)
}
