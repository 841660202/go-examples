package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	fmt.Println(os.Args)
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
