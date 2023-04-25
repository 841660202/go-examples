package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Boottime String `json:"boottime"` //这里是我们接收时自定义的参数类型String
}
type String string

func main() {
	app := &App{}
	a := `{"boottime":1732323232}`
	err := json.Unmarshal([]byte(a), &app)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(app)
}

// 这里做类型转换
func (s *String) UnmarshalJSON(data []byte) error {
	*s = String(data)
	fmt.Println("*s", *s)
	return nil
}
