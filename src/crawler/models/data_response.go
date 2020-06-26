package models

import "github.com/jinzhu/gorm"

type DataCrawl struct {
	gorm.Model
	ResponseBody string
	TimeExec float32
}
