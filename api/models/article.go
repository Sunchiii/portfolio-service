package models

import (
	"encoding/json"
	"time"
)

type Article struct {
	ID          int64           `json:"id"`
	UserId      int64           `json:"user_id"`
	ImageExam   string          `json:"image_exam"`
	ArticleType string          `json:"article_type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Data        json.RawMessage `json:"data"`
	CreatedAt   time.Time       `json:"created_at"`
}
