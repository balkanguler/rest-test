package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()

	router.POST("/concats", func(c *gin.Context) {

		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
	
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}