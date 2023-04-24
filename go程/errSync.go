package main

import (
	"fmt"
	"time"
)

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(",")
	print(a)
}

func main() {
	go f()
	g()
	time.Sleep(time.Second)
	fmt.Println("\n---------------")
	g()
}
