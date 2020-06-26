package jobs

import "fmt"

type JobCrawler struct {
	Url string
}

func (j JobCrawler) PushQueue()  {
	fmt.Println("push queue crawl " + j.Url)
	CrawlerQueue <- j
}

var CrawlerQueue = make(chan JobCrawler, 100)