package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunchiii/portfolio-service/api/models"
	"github.com/sunchiii/portfolio-service/pkg/database"
	"github.com/sunchiii/portfolio-service/pkg/utils"
)

type ArticleHandler struct {
	Db *database.DB
}

func NewArticleHandler(db *database.DB) (ArticleHandler, error) {
	return ArticleHandler{Db: db}, nil
}

func (articledb ArticleHandler) CreateAtricle(c *gin.Context) {
	// prepare article object
	var article models.Article

	// bind json data from user to article object
	if err := c.ShouldBindJSON(&article); err != nil {
		errMsg := utils.BadRequestError("check your data request")
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	// prepare newdata of article
	newArticle := models.Article{
		ID:          int64(uuid.New().ID()),
		Title:       article.Title,
		Description: article.Description,
		Data:        article.Data,
		CreatedAt:   time.Now(),
	}

	// call function database ti create new article
	err := articledb.Db.CreateArticle(&newArticle)
	if err != nil {
		errMsg := utils.InternalServerError("something wrong in server side")
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}

func (articledb ArticleHandler) GetArticles(c *gin.Context) {
	// call database
	article, err := articledb.Db.GetArticles(1, 10)
	if err != nil {
		errMsg := utils.InternalServerError("can't get data from database")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	c.JSON(http.StatusOK, article)
}

func (articledb ArticleHandler) GetArticle(c *gin.Context) {
	// Get the user ID from the request parameters
	articleID := c.Param("id")
	// Retrieve the user from the database or any other data source

	article, err := articledb.Db.GetArticle(articleID)
	if err != nil {
		errMsg := utils.InternalServerError("can't query data from database")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	// Return the user as JSON response
	c.JSON(http.StatusOK, article)
}


func (articledb ArticleHandler) UpdateArticle(c *gin.Context) {
	// get id from param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg := utils.BadRequestError("id should be number only")
    log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
    return
	}

	// Parse the request body to get the user data
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user data
	oldArticle, err := articledb.Db.GetArticle(strconv.Itoa(id))
	if err != nil {
		errMsg := utils.BadRequestError("does't exit")
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	// prepare data befor insert to database
	newArticle := models.Article{
		ID:        int64(id),
    Title: article.Title,
    Description: article.Description,
    Data: article.Data,
    CreatedAt: time.Now(),
	}

	if article.Title == "" {
		newArticle.Title = oldArticle.Title
	}
	if article.Description == "" {
		newArticle.Description = oldArticle.Description
	}
	if len(article.Data) <=0 {
		newArticle.Data = oldArticle.Data
	}

	// insert data to database
	err = articledb.Db.UpdateArticle(&newArticle)
	if err != nil {
		errMsg := utils.InternalServerError("can't update data to database")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg)
		return
	}
	// Return a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

