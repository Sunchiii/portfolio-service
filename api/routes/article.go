package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/handler"
	"github.com/sunchiii/portfolio-service/pkg/database"
)

func ArticleRoutes(route *gin.Engine,db *database.DB){
  articlesHandler,err := handler.NewArticleHandler(db)
  if err != nil{
    log.Fatal("can't connect handler",err)
  }
  v1 := route.Group("/v1")
  {
    v1.GET("/articles",articlesHandler.GetArticles)
    v1.GET("/article/:id",articlesHandler.GetArticle)
    v1.POST("/article",articlesHandler.CreateAtricle)
  }
}




