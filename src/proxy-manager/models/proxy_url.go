package models

import (
	"github.com/jinzhu/gorm"
	"go.opentelemetry.io/otel/api/trace"
	"golang.org/x/net/html/atom"
)

type ProxyUrl struct {
	gorm.Model
	ProxyID uint `gorm:"unique_index:idx_proxy_url"`
	Proxy []Proxy
	Url []Url
	UrlID uint `gorm:"unique_index:idx_proxy_url"`
	NumberRequestSuccess int
	NumberRequestFail int
}

func ChooseProxy(urlString string) Proxy {
	proxyResult := Proxy{}

	return proxyResult
}

func Save(url Url, proxy Proxy, db *gorm.DB) {
	
}

func RecordRequest(url Url, proxy Proxy, db *gorm.DB, isSuccess bool) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		row := tx.Where(&ProxyUrl{
			ProxyID: proxy.ID,
			UrlID:   url.ID,
		})
		if isSuccess == true {
			if err:= row.Update("number_request_success", gorm.Expr("number_request_success + ?", 1)).Error; err != nil {
				return err
			}
		} else {
			if err:= row.Update("number_request_fail", gorm.Expr("number_request_fail + ?", 1)).Error; err != nil {
				return err
			}
		}

		row = tx.Where(&Proxy{})

		if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil{
			// return any error will rollback
			return err
		}

		if err != nil {
			return err
		}

		// return nil will commit
		return nil
	})
}