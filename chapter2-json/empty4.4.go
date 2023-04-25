package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password,-"`
}

type PublicUser struct {
	*User // 匿名嵌套
	// Password *struct{} `json:"password,omitempty"`
}

func omitPasswordDemo() {
	u1 := User{
		Name:     "左右逢源",
		Password: "123456",
	}

	u2 := PublicUser{User: &u1}
	fmt.Printf("stu: %v\n", u2)      // stu: {0xc0000b6000 <nil>}
	fmt.Printf("stu: %v\n", u2.User) // stu: &{左右逢源 123456}

	b, err := json.Marshal(PublicUser{User: &u1}) // 这里在初始化值的时候

	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}

	fmt.Printf("str:%s\n", b) // str:{"name":"左右逢源"}
}

func main() {
	omitPasswordDemo()
}

// 不修改原结构体忽略空值字段
// 我们需要json序列化User，但是不想把密码也序列化，又不想修改User结构体，
// 这个时候我们就可以使用创建另外一个结构体PublicUser匿名嵌套原User，
// 同时指定Password字段为匿名结构体指针类型，并添加omitemptytag，示例代码如下：
