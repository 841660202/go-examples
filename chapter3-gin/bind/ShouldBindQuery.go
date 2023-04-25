package main

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Age     *int   `form:"age"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println("相等吗:", person.Name)
		log.Println("Address:", person.Address)
		log.Printf("person: %+#v", person)
		log.Printf("person.Name: %v", person.Name)
		log.Printf("person.Age: %t", person.Age == nil)
		log.Printf("person.Age: %t", IsNil(person.Age))
	}
	c.String(200, "Success")
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// https://www.cnblogs.com/rainbow-tan/p/15457818.html

// golang的零值问题如何处理？

// 1. 假设是指针，指针需要判空，转化成gorm查询数据库，又要组装sql，还有可能需要去掉零值/空值
