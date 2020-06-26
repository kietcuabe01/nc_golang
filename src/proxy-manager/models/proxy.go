package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type Proxy struct {
	gorm.Model
	IpPort      string `gorm:"type:varchar(100);unique_index"`
	IsAvailable bool
	AvailableAt time.Time
	UserID      int
	User		User
	ProxyUrl    []ProxyUrl
}

func GetProxyFree(db *gorm.DB) []Proxy  {
	var proxies []Proxy

	db.Where("`user_id` = ?", 0).Find(&proxies)

	return proxies
}

func ReleaseIp(ip string, db *gorm.DB) Proxy {
	proxy := Proxy{}
	db.Where(&Proxy{IpPort: ip}).Update(&Proxy{UserID: 0})

	return proxy
}