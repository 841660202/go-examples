package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Hobby []string `json:"hobby"`
}

func omitemptyDemo() {
	u1 := User{
		Name: "左右逢源",
	}
	// struct -> json string
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// str:{"name":"左右逢源","email":"","hobby":null}
}

func main() {
	omitemptyDemo()
}
