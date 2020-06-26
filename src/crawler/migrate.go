package main

import (
	"crawler/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main()  {
	db, error := gorm.Open("mysql", "remote:Remote@123456@tcp(127.0.0.1:3306)/go_crawler?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	fmt.Print(error)
	db.AutoMigrate(&models.DataCrawl{})

}