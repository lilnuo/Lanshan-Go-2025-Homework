package main

import "fmt"

func add(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func main() {
	var num int
	fmt.Scanf("%d", &num)
	result := add(num)
	fmt.Println(result)

}
