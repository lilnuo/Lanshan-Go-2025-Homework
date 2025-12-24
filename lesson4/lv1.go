package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func writeWithoutBuffer(filename string, data []byte, times int) time.Duration {
	start := time.Now()
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := 0; i < times; i++ {
		_, err := file.Write(data)
		if err != nil {
			panic(err)
		}
	}
	return time.Since(start)
}
func writeWithBuffer(filename string, data []byte, times int) time.Duration {
	start := time.Now()
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < times; i++ {
		_, err := writer.Write(data)
		if err != nil {
			panic(err)
		}
	}
	return time.Since(start)
}

func main() {
	testDate := []byte("Hello,this is test data for I/O performance comparison!\n")
	writeCounts := []int{1000, 10000, 100000, 1000000}
	fmt.Println("I/O performance comparison")
	fmt.Println("=============================")
	for _, count := range writeCounts {
		fmt.Printf("Write count:%d\n", count)
		fmt.Println("---------------------")
		file1 := fmt.Sprintf("test_no_buffer_%d.txt", count)
		duration1 := writeWithoutBuffer(file1, testDate, count)
		fmt.Printf("不带缓冲:%v\n", duration1)

		file2 := fmt.Sprintf("test_buffer_%d.txt", count)
		duration2 := writeWithBuffer(file2, testDate, count)
		fmt.Printf("带缓冲:%v\n", duration2)

		improvement := float64(duration1-duration2) / float64(duration1) * 100
		fmt.Printf("性能提升:%v\n", improvement)
		os.Remove(file1)
		os.Remove(file2)

	}
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("大数据详细对此（100，000次写入）")
	fmt.Println(strings.Repeat("=", 50))

	largeCount := 10
	largeData := []byte("this is a large data string to demonstrate the performance difference more clearly,\n")
	var totalWithoutBuffer, totalWithBuffer time.Duration
	testRounds := 5
	for i := 0; i < testRounds; i++ {
		file1 := fmt.Sprintf("test_no_buffer_%d.txt", i)
		file2 := fmt.Sprintf("test_buffer_%d.txt", i)
		totalWithoutBuffer += writeWithoutBuffer(file1, largeData, largeCount)
		totalWithBuffer += writeWithBuffer(file2, largeData, largeCount)
		os.Remove(file1)
		err := os.Remove(file2)
		if err != nil {
			return
		}
	}
	avgWithoutBuffer := totalWithoutBuffer / time.Duration(testRounds)
	avgWithBuffer := totalWithBuffer / time.Duration(testRounds)
	avgImprovement := float64(avgWithoutBuffer-avgWithBuffer) / float64(avgWithoutBuffer) * 100
	fmt.Printf("平均耗时-不带缓冲：%v\n", avgWithoutBuffer)
	fmt.Printf("平均耗时-带缓冲：%v\n", avgWithBuffer)
	fmt.Printf("平均性能提升：%.2f%%\n", avgImprovement)

}
