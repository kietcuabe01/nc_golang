package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main()  {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		query := c.Query("query")
		r := rand.Intn(10)
		sleep := time.Duration(r) * time.Millisecond * 1000
		time.Sleep(sleep)
		c.JSON(200, gin.H{
			"time" : r,
			"message": "message_" + query,
		})
	})

	route.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
