package worker

import (
	"crawler/database"
	"crawler/jobs"
	"crawler/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

// Todo struct
type ResponseBodyCrawled struct {
	Message string `json:"message"`
	Time  float32 `json:"time"`
}

func Crawl (url string) (ResponseBodyCrawled, error) {
	responseStruct := ResponseBodyCrawled{Message: "", Time: 0}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return responseStruct, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to struct
	_ = json.Unmarshal(bodyBytes, &responseStruct)
	return responseStruct, nil
}

func WorkCrawl(wg *sync.WaitGroup) {
	fmt.Println("begin WorkCrawl")
	for job := range jobs.CrawlerQueue {
		response, _ := Crawl(job.Url)
		fmt.Println("do WorkCrawl " + job.Url)
		bytes, _ := json.Marshal(response)
		result := jobs.ResultCrawler{
			Response: string(bytes),
			Time: response.Time,
		}
		jobs.ResultCrawler.PushQueue(result)
	}
	wg.Done()
}

func CreateWorkerCrawlPool(noOfWorkers int) {
	fmt.Println("CreateWorkerCrawlPool " + strconv.Itoa(noOfWorkers))
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go WorkCrawl(&wg)
	}
	wg.Wait()
	close(jobs.ResultQueue)
}

//--------

func WorkSaveDatabase(wg *sync.WaitGroup, db *gorm.DB) {
	fmt.Println("begin WorkSaveDatabase")
	for job := range jobs.ResultQueue {
		model := models.DataCrawl{ResponseBody: job.Response, TimeExec: job.Time}
		fmt.Println("do Work Save dataabse " + job.Response)
		database.SaveResponse(model, db)
	}
	wg.Done()
}

func CreateWorkerDatabasePool(noOfWorkers int, db *gorm.DB) {
	fmt.Println("CreateWorkerDatabasePool " + strconv.Itoa(noOfWorkers))
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go WorkSaveDatabase(&wg, db)
	}
	wg.Wait()
}