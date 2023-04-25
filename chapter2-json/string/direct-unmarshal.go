package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`

	// 我看直接解析成数组也没什么问题

	var arr []interface{}

	if err := json.Unmarshal([]byte(blob), &arr); err != nil {
		log.Fatal(err)
	}

	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, " : ", v)
	}

}
