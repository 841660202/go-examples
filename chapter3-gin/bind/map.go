package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
		fmt.Println("--------------------------------")
		fmt.Printf("ids: %+v; names: %+v", ids, names)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
	router.Run(":8080")
}

// curl --location -g --request POST 'http://localhost:8080/post?ids[a]=1234&ids[b]=hello'
// curl --location -g --request POST 'http://localhost:8080/post?names[first]=thinkerou&names[second]=tianou'

// 接收前端传递的对象，转化成map
