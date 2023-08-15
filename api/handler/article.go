package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (articledb *ArticleHandler) CreateAtricle(c *gin.Context) {
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
		Title:       article.Title,
		Description: article.Description,
		Data:        article.Data,
    UserId: article.UserId,
    ImageExam: article.ImageExam,
    ArticleType: article.ArticleType,
	}

	// call function database ti create new article
	err := articledb.Db.CreateArticle(&newArticle)
	if err != nil {
		errMsg := utils.InternalServerError("something wrong in server side")
    log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}

func (articledb *ArticleHandler) GetArticles(c *gin.Context) {
  page := c.DefaultQuery("page","1")
  limit := c.DefaultQuery("limit","10")

  // convert page and limit to int
  pageValue, _ := strconv.Atoi(page)
  limitValue, _ := strconv.Atoi(limit)
	// call database
	article, err := articledb.Db.GetArticles(pageValue, limitValue)
	if err != nil {
		errMsg := utils.InternalServerError("can't get data from database")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	c.JSON(http.StatusOK, article)
}

func (articledb *ArticleHandler) GetArticle(c *gin.Context) {
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


func (articledb *ArticleHandler) UpdateArticle(c *gin.Context) {
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
    log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user data
	oldArticle, err := articledb.Db.GetArticle(strconv.Itoa(id))
	if err != nil {
		errMsg := utils.BadRequestError("does't exit")
    log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}


	// prepare data befor insert to database
	newArticle := models.Article{
    Title: article.Title,
    Description: article.Description,
    Data: article.Data,
    ImageExam: article.ImageExam,
    ArticleType: article.ArticleType,
    UserId: oldArticle.UserId,
    ID: oldArticle.ID,
	}

	if article.Title == "" {
		newArticle.Title = oldArticle.Title
	}
	if article.Description == "" {
		newArticle.Description = oldArticle.Description
	}
  if article.ImageExam == ""{
    newArticle.ImageExam = oldArticle.ImageExam
  }
  if article.ArticleType == ""{
    newArticle.ArticleType = oldArticle.ArticleType
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
	c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})
}



func (articledb *ArticleHandler) DeleteArticle(c *gin.Context) {
	// Get the user ID from the request parameters
	articleID := c.Param("id")

	// call database
	err := articledb.Db.DeleteArticle(articleID)
	if err != nil {
		errMsg := utils.InternalServerError("something wrong in server")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "article deleted successfully"})
}

