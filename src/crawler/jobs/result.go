package jobs

type ResultCrawler struct {
	Response string
	Time float32
}

func (r ResultCrawler) PushQueue ()  {
	ResultQueue <- r
}

var ResultQueue = make(chan ResultCrawler, 100)