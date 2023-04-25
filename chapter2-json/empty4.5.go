package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"-"`
	Money    int    `json:"-,"`
	Skip     int    `json:",omitempty"`
	Clothing int    `json:"clothing,-"`
}

func main() {

	user := User{
		ID:       1,
		Name:     "John",
		Age:      30,  // 被忽略
		Money:    100, // 键被改为-
		Clothing: 10,  // -没有生效
		Skip:     3,
	}

	b, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}

// 总结
// Field is ignored by this package.
// Field int `json:"-"`

// Field appears in JSON as key "-".
// Field int `json:"-,"`
