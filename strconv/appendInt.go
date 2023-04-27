package main

import (
	"fmt"
	"strconv"
)

func main() {
	var dst []byte
	num := int64(42)
	base := 2
	fmt.Println(dst)
	dst = strconv.AppendInt(dst, num, base)
	fmt.Println(dst)
	fmt.Println(string(dst))
}
