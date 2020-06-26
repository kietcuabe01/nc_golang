package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"proxy-manager/models"
	"time"
)

func main()  {
	db, error := gorm.Open("mysql", "remote:Remote@123456@tcp(127.0.0.1:3306)/golang_proxy_manager?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if error != nil {
		fmt.Printf("Erorr %s", error)
	}
	db.AutoMigrate(&models.User{}, &models.Proxy{}, &models.Url{}, &models.ProxyUrl{})

	user := models.User{}
	db.Where("`key` = ?", "123456789").First(&user)
	if user.Email == "" {
		user = models.User{
			Name:  "Test",
			Email: "kiet.tran@epsilo.io",
			Key:   "123456789",
		}

		db.Create(&user)
	}

	proxy1 := models.Proxy{}
	proxy1.IpPort = "0.0.0.0:8111"
	proxy1.IsAvailable = true
	proxy1.AvailableAt = time.Now()
	db.Create(&proxy1)

	proxy2 := models.Proxy{}
	proxy2.IpPort = "0.0.0.0:8222"
	proxy2.IsAvailable = true
	proxy2.AvailableAt = time.Now()
	db.Create(&proxy2)

	proxy3 := models.Proxy{}
	proxy3.IpPort = "0.0.0.0:8333"
	proxy3.IsAvailable = true
	proxy3.AvailableAt = time.Now()
	db.Create(&proxy3)

}
