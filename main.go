package main

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/gin-contrib/cors"

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/concats", func(c *gin.Context) {

		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
	
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}