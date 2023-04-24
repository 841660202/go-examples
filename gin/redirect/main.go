package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	router.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	router.GET("/hello", func(c *gin.Context) {
		fmt.Println("hello")

		c.Set("name", "john") // set name=john

		c.Request.URL.Path = "/hi"
		router.HandleContext(c)
	})

	router.GET("/hi", func(c *gin.Context) {
		fmt.Println("hi")

		name, _ := c.Get("name") // name is nil

		c.JSON(200, gin.H{"message": name})
	})

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}

// $ curl http://localhost:8080/test
// <a href="http://www.google.com/">Moved Permanently</a>.

// HandleContext 内部会执行重置
// https://matthung0807.blogspot.com/2021/11/gin-engine-handlercontext-func.html
