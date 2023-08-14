package database

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sunchiii/portfolio-service/api/models"
)

type DB struct {
	*sql.DB
}

type ArticlesResponse struct {
	Articles []*models.Article `json:"articles"`
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

func (db *DB) CreateUser(user *models.User) error {
	sqlStatement := `
		INSERT INTO "user" (id,username, password,created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	err := db.QueryRow(sqlStatement,user.ID, user.Username, user.Password,user.CreatedAt).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateUser(user *models.User) error{
  sqlStatement := `UPDATE "user" SET username = $2, password = $3 WHERE id = $1`
  // prepare sql statement
  stmt,err := db.Prepare(sqlStatement)
  if err != nil{
    return err
  }
  defer stmt.Close()

  // execute statement 
  _, err = stmt.Exec(user.ID,user.Username,user.Password)
  if err != nil{
    return err
  }
  return nil
}

func (db *DB) GetUsers() ([]*models.User, error) {
	sqlStatement := `SELECT id, username, password, created_at FROM "user"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		var user models.User
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
func (db *DB) GetUser(_id string) (*models.User, error) {
  // convert id string to int
  id,err := strconv.Atoi(_id)
  if err != nil{
    return nil,err
  }
  var user models.User
  // Prepare the SQL statement
  stmt,err := db.Prepare(`SELECT id, username,password,created_at FROM "user" WHERE id = $1`)
  if err != nil{
    log.Println(err)
    return nil,err
  }
  defer stmt.Close()

  // Execute the query and retrieve the user data
  err = stmt.QueryRow(id).Scan(&user.ID, &user.Username,&user.Password, &user.CreatedAt)
	if err != nil {
    log.Println("can't execute command")
		return nil, err
	}

  return &user,nil
}


func (db *DB) GetUserByUsername(_username , _password string) (*models.User, error) {
  var user models.User
  // Prepare the SQL statement
  sqlStatement := `SELECT id, username, password,created_at FROM "user" WHERE username = $1 AND password = $2`

  // Execute the query and retrieve the user data
  row := db.QueryRow(sqlStatement,_username,_password)

  // recieve the user data
  if err := row.Scan(&user.ID,&user.Username,&user.Password,&user.CreatedAt); err != nil{
    return nil,err
  }

  return &user,nil
}


func (db *DB) DeleteUser(_id string) error{
  // Prepare a SQL statement to delete the user with the given ID
  stmt, err := db.Prepare(`DELETE FROM "user" WHERE id = $1`)
  if err != nil {
    return err
  }
  defer stmt.Close()
  
  // Execute the statement with the ID parameter
  _, err = stmt.Exec(_id)
  if err != nil {
    return err
  }
  return nil
}

func (db *DB) CreateArticle(article *models.Article) error {
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

	articles := []*models.Article{}
	for rows.Next() {
		var article models.Article
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


func (db *DB) GetArticle(_id string) (*models.Article, error) {
  var article models.Article
  // Prepare the SQL statement
  stmt,err := db.Prepare(`SELECT title, description,data,created_at FROM "article" WHERE id = $1`)
  if err != nil{
    log.Println(err)
    return nil,err
  }
  defer stmt.Close()

  // Execute the query and retrieve the user data
  err = stmt.QueryRow(_id).Scan(&article.Title,&article.Description,&article.Data,&article.CreatedAt)
	if err != nil {
    log.Println(err)
		return nil, err
	}

  return &article,nil
}

func (db *DB) UpdateArticle(article *models.Article) error{
  sqlStatement := `UPDATE "article" SET title = $2, description = $3, data = $4 WHERE id = $1`
  // prepare sql statement
  stmt,err := db.Prepare(sqlStatement)
  if err != nil{
    return err
  }
  defer stmt.Close()

  // execute statement 
  _, err = stmt.Exec(article.ID, article.Title, article.Description, article.Data)
  if err != nil{
    return err
  }
  return nil
}

func (db *DB) DeleteArticle(_id string) error{
  // Prepare a SQL statement to delete the user with the given ID
  stmt, err := db.Prepare(`DELETE FROM "article" WHERE id = $1`)
  if err != nil {
    return err
  }
  defer stmt.Close()
  
  // Execute the statement with the ID parameter
  _, err = stmt.Exec(_id)
  if err != nil {
    return err
  }
  return nil
}

