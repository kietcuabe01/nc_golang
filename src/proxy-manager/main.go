package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"proxy-manager/models"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	db, _ := gorm.Open("mysql", "remote:Remote@123456@tcp(127.0.0.1:3306)/golang_proxy_manager?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	route := gin.Default()
	ctx := context.Background()

	route.GET("/get", func(c *gin.Context) {
		keyUser := c.Query("key")
		urlString := c.Query("url")
		//urlDB := models.FindCreate(urlString, db)
		user := models.GetUser(client, keyUser, ctx, db)
		if user.Email != "" {
			proxyResponse := models.ChooseProxy(urlString)

			c.JSON(200, gin.H{
				"message": "success",
				"ip" : proxyResponse.IpPort,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "User not found",
			})
		}
	})

	route.POST("/release", func(c *gin.Context) {
		keyUser := c.Query("key")
		ip := c.Query("ip_port")
		isSucess := c.Query("is_success")
		user := models.GetUser(client, keyUser, ctx, db)
		if user.Email != "" {

			c.JSON(200, gin.H{
				"message": "User " + user.Email,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "User not found",
			})
		}
	})

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}