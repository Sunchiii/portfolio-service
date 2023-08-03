package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/handler"
	"github.com/sunchiii/portfolio-service/pkg/database"
)

func UserRoutes(route *gin.Engine,db *database.DB){
  usersHandler,err := handler.NewUserHandler(db)
  if err != nil{
    log.Fatal("can't connect handler",err)
  }
  v1 := route.Group("/v1")
  {
    v1.GET("/users",usersHandler.GetUsersHandler)
  }
}


