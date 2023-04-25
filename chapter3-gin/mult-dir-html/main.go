package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}

// 测试

// $ curl http://localhost:8080/posts/index

// <html><h1>
//         Posts
// </h1>
// <p>Using posts/index.tmpl</p>
// </html>

// curl http://localhost:8080/users/index

// <html><h1>
//         Users
// </h1>
// <p>Using users/index.tmpl</p>
// </html>
