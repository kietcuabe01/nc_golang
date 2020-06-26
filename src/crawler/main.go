package main

import (
	"crawler/database"
	"crawler/jobs"
	"crawler/worker"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func allocate(urls []string)  {
	for _, url := range urls {
		job := jobs.JobCrawler{
			Url: url,
		}
		job.PushQueue()
	}
	close(jobs.CrawlerQueue)
}

func main() {
	startTime := time.Now()
	db, _ := gorm.Open("mysql", "remote:Remote@123456@tcp(127.0.0.1:3306)/go_crawler?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	urls := database.GetUrls()
	go allocate(urls)
	go worker.CreateWorkerCrawlPool(900)
	worker.CreateWorkerDatabasePool(100, db)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}