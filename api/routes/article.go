package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/handler"
	"github.com/sunchiii/portfolio-service/api/middleware"
	"github.com/sunchiii/portfolio-service/pkg/database"
)

func ArticleRoutes(route *gin.Engine,db *database.DB){
  articlesHandler,err := handler.NewArticleHandler(db)
  if err != nil{
    log.Fatal("can't connect handler",err)
  }
  v1x := route.Group("/v1")
  v1y := route.Group("/v1")
  v1y.Use(middleware.AuthMidleware())
  {
    v1x.GET("/articles",articlesHandler.GetArticles)
    v1x.GET("/article/:id",articlesHandler.GetArticle)
  }

  {
    v1y.POST("/article",articlesHandler.CreateAtricle)
    v1y.PUT("/article/:id", articlesHandler.UpdateArticle)
    v1y.DELETE("/article/:id", articlesHandler.DeleteArticle)
  }
}




