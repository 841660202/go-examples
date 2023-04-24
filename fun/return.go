package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// 非命名返回
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// https://learnku.com/docs/the-way-to-go/function-parameters-and-return-values/3600#e1892d
