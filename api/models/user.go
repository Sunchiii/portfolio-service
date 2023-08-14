package models

import (
  "time"
)

type User struct {
	ID        int64     `json:"id"`
  ArticleId int64 `json:"article_type"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

