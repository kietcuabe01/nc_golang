package database

import (
	"crawler/models"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func GetUrls() []string {
	var urls []string
	for i := 0; i < 100000; i++ {
		urls = append(urls, "http://0.0.0.0:8080?query=" + strconv.Itoa(i))
	}

	return urls
}

func SaveResponse(response models.DataCrawl, db *gorm.DB) {
	time.Sleep(500 * time.Millisecond)
	db.Create(&response)
}