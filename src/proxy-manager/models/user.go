package models

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Num int `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	Name string
	Email string `gorm:"type:varchar(100);unique_index"`
	Key string
}

func GetUser(client *redis.Client, keyUser string, ctx context.Context, db *gorm.DB) User {
	var user = User{}
	if keyUser == "" {
		return user
	}

	resultCache, errorCache := client.Get(ctx, "proxy-manager:user-" + keyUser).Result()

	if errorCache == redis.Nil {
		db.Where("`key` = ?", keyUser).First(&user)
		bytes, _ := json.Marshal(user)
		client.Set(ctx, "proxy-manager:user-" + keyUser, bytes, 0)
	} else {
		_ = json.Unmarshal([]byte(resultCache), &user)
	}

	return user
}