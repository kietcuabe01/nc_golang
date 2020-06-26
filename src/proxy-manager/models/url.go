package models

import (
	"github.com/jinzhu/gorm"
)

type Url struct {
	gorm.Model
	Url    string
	ProxyUrl  []ProxyUrl
}

func FindCreate(url string, db *gorm.DB) Url {
	urlDb := Url{Url: url}
	db.Where(Url{Url: url}).FirstOrCreate(&urlDb)
	return urlDb
}