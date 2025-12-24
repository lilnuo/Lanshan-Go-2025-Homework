package main

import "fmt"

func average(sum int, count int) float64 {
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
func main() {
	var num, sum, count int
	fmt.Println("请输入一个整数（输入0结束）：")
	for {
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		sum += num
		count++
	}
	avg := average(sum, count)
	if count == 0 {
		fmt.Println("没有任何成绩")
		return
	} else if avg >= 60 {
		fmt.Println("成绩为%.2f,成绩合格\n", avg)
	} else {
		fmt.Println("平均成绩为%.2f,成绩不合格\n", avg)
	}

}
