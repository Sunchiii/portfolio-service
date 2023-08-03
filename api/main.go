package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/routes"
	"github.com/sunchiii/portfolio-service/config"
	"github.com/sunchiii/portfolio-service/pkg/database"
)

func main(){
  // initial config
  newConfig,err := config.NewConfig()
  if err != nil{
    log.Fatal("can't initial without config")
  }

  //initial database with our config
  db,err := database.NewDB(newConfig.PGUrl)
  if err != nil{
    log.Fatal("can't connect to database!!",err)
  }
  // initial ginEngin
  r := gin.Default()

  routes.UserRoutes(r,db)
  r.Run(":"+newConfig.Port) // listen and serve on 0.0.0.0:8080
}
