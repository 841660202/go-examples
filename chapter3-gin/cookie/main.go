package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"设置": "cookie",
			})
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}

// curl -v http://localhost:8080/cookie
