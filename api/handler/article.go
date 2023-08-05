package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunchiii/portfolio-service/api/models"
	"github.com/sunchiii/portfolio-service/pkg/database"
	"github.com/sunchiii/portfolio-service/pkg/utils"
)

type ArticleHandler struct{
  Db *database.DB
}

func NewArticleHandler(db *database.DB)(ArticleHandler,error){
  return ArticleHandler{Db: db},nil
}

func (articledb ArticleHandler) CreateAtricle(c *gin.Context){
  // prepare article object
  var article models.Article 

  // bind json data from user to article object
	if err := c.ShouldBindJSON(&article); err != nil {
    errMsg := utils.BadRequestError("check your data request")
    c.JSON(errMsg.Status,errMsg.Message)
		return
	}

  // prepare newdata of article
  newArticle := models.Article{
    ID: int64(uuid.New().ID()),
    Title: article.Title,
    Description: article.Description,
    Data: article.Data,
    CreatedAt: time.Now(),
  }

  // call function database ti create new article
  err := articledb.Db.CreateArticle(&newArticle)
  if err != nil{
    errMsg := utils.InternalServerError("something wrong in server side")
    c.JSON(errMsg.Status,errMsg.Message)
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "Article created successfully"})

}
