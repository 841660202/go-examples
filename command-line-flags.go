package main

import (
	"flag"
	"fmt"
)

// 和nodejs开发命令类似的东西，指定名称的参数/默认参数，写法不一样罢了
func main() {

	// String defines a string flag with specified name, default value, and usage string.
	wordPtr := flag.String("word", "foo", "a string")
	// Int defines an int flag with specified name, default value, and usage string.
	// The return value is the address of an int variable that stores the value of the flag.
	numbPtr := flag.Int("numb", 42, "an int")
	// Bool defines a bool flag with specified name, default value, and usage string.
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
