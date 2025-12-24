package main

import "fmt"

func main() {
	i := 1
	sum := 0
	for i = 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}
