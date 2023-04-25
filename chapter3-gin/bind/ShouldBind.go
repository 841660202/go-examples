package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
	// Timestamp time.Time `form:"timestamp"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	if err := c.ShouldBind(&person); err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else {
		fmt.Printf("%v", err)
		log.Print(err)
	}

	c.String(200, "Success")
}

// $ curl -X GET "localhost:8080/testing?name=appleboy&address=xyz&birthday=1992-03-15"
