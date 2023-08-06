package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/handler"
	"github.com/sunchiii/portfolio-service/api/middleware"
	"github.com/sunchiii/portfolio-service/pkg/database"
)

func UserRoutes(route *gin.Engine,db *database.DB){
  usersHandler,err := handler.NewUserHandler(db)
  if err != nil{
    log.Fatal("can't connect handler",err)
  }
  v1 := route.Group("/v1")
  v1.Use(middleware.AuthMidleware())
  {
    v1.GET("/users",usersHandler.GetUsersHandler)
    v1.GET("/user/:id",usersHandler.GetUserHandler)
    v1.POST("/user",usersHandler.CreateUserHandler)
    v1.PUT("/user/:id",usersHandler.UpdateUserHandler)
    v1.DELETE("/user/:id",usersHandler.DeleteUserHandler)
  }
}


