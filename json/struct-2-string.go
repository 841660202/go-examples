package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string //如果改成小写，序列化时跨包会丢掉此字段
	Age   int    `json:"age"` //代表json过后的重命名
	Score float64
}

func testStruct() {
	monster := Monster{Name: "牛魔王", Age: 300, Score: 95}
	data, err := json.Marshal(&monster) //返回的是[]byte切片
	if err != nil {
		fmt.Printf("序列化错误err=%v", err)
	}
	fmt.Printf("序列化结构体=%v\n", string(data))
}

func testSlice() {
	var monster []map[string]interface{}
	a := make(map[string]interface{})
	a["name"] = "狮子精"
	a["age"] = "320"
	a["score"] = 92.5

	b := make(map[string]interface{})
	b["name"] = "大象精"
	b["age"] = "560"
	b["score"] = 90

	monster = append(monster, a, b)

	data, err := json.Marshal(monster) //返回的是[]byte切片
	if err != nil {
		fmt.Printf("序列化错误err=%v", err)
	}
	fmt.Printf("序列化切片=%v\n", string(data))
}

func testMap() {
	var monster map[string]interface{}
	monster = make(map[string]interface{})
	monster["name"] = "白骨精"
	monster["age"] = 300
	monster["score"] = 95.5

	data, err := json.Marshal(monster) //返回的是[]byte切片
	if err != nil {
		fmt.Printf("序列化错误err=%v", err)
	}
	fmt.Printf("序列化map=%v\n", string(data))

}

// json.Marshal 将数据编码成json字符串
func main() {
	fmt.Println("\n序列化结构体----------------------------------------")
	testStruct() //序列化结构体

	fmt.Println("\n序列化map----------------------------------------")
	testMap() //序列化map

	fmt.Println("\n序列化切片----------------------------------------")
	testSlice() //序列化切片

	fmt.Println("\n反序列化struct----------------------------------------")
	unmarshalStruct() //反序列化struct

	fmt.Println("\n反序列化map----------------------------------------")
	unmarshalMap() //反序列化map

	fmt.Println("")
	fmt.Println("\n反序列化slice----------------------------------------")
	unmarshalSlice() //反序列化slice
}

func unmarshalStruct() {
	var monster Monster
	str := "{\"Name\":\"牛牛\",\"age\":330,\"Score\":92}"
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后：%v\n", monster)
}

func unmarshalMap() {
	var a map[string]interface{}
	str := "{\"age\":300,\"name\":\"白骨精\",\"score\":95.5}"
	//反序列化map，不需要make,make操作被封装到Unmarshal函数中
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后：%v", a)
}

func unmarshalSlice() {
	var a []map[string]interface{}
	str := "[{\"age\":\"320\",\"name\":\"狮子精\",\"score\":92.5},{\"age\":\"560\",\"name\":\"大象精\",\"score\":90}]"
	//反序列化map，不需要make,make操作被封装到Unmarshal函数中
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}
	fmt.Printf("反序列化后：%v", a)
}
