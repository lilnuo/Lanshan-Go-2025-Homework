package main

import "fmt"

func main() {
	testArray := [5]int{1, 2, 3, 4, 5}
	output := ab(testArray)
	fmt.Println("元素出现次数：")
	for num, count := range output {
		fmt.Printf("%d: %d次\n", num, count)
	}
	fmt.Println(output)

}

func ab(arr [5]int) map[int]int {
	result := make(map[int]int)
	for _, num := range arr {
		result[num]++
		//返回map
	}
	return result
}
