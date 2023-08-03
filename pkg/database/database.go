package database

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Data        json.RawMessage `json:"data"`
	CreatedAt   time.Time       `json:"created_at"`
}

type DB struct {
	*sql.DB
}

type ArticlesResponse struct {
	Articles []*Article `json:"articles"`
	Meta     Meta       `json:"meta"`
}

type Meta struct {
	TotalCount int `json:"total_count"`
	PageCount  int `json:"page_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) CreateUser(user *User) error {
	sqlStatement := `
		INSERT INTO "user" (username, password, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`
	err := db.QueryRow(sqlStatement, user.Username, user.Password, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetUsers() ([]*User, error) {
	sqlStatement := `SELECT id, username, password, created_at FROM "user"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *DB) CreateArticle(article *Article) error {
	sqlStatement := `
		INSERT INTO "article" (title, description, data, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	err := db.QueryRow(sqlStatement, article.Title, article.Description, article.Data, article.CreatedAt).Scan(&article.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetArticles(page int, perPage int) (*ArticlesResponse, error) {
	// Get total count of articles
	sqlStatement := `SELECT COUNT(*) FROM "article"`
	var totalCount int
	err := db.QueryRow(sqlStatement).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// Calculate pagination metadata
	pageCount := (totalCount + perPage - 1) / perPage
	offset := (page - 1) * perPage

	// Get articles for the requested page
	sqlStatement = `
		SELECT id, title, description, data, created_at
		FROM "article"
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	rows, err := db.Query(sqlStatement, perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*Article{}
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.Data, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	meta := Meta{
		TotalCount: totalCount,
		PageCount:  pageCount,
		Page:       page,
		PerPage:    perPage,
	}

	response := &ArticlesResponse{
		Articles: articles,
		Meta:     meta,
	}

	return response, nil
}
