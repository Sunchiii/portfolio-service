package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	stdmdw "github.com/sunchiii/portfolio-service/api/middleware"
	"github.com/sunchiii/portfolio-service/api/routes"
	"github.com/sunchiii/portfolio-service/config"
	"github.com/sunchiii/portfolio-service/pkg/database"
  limiter "github.com/ulule/limiter/v3"
  mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
  "github.com/ulule/limiter/v3/drivers/store/memory"

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

  // Define a limit rate to times requests per second.
  rate := limiter.Rate{
    Period: 1 * time.Second,
    Limit:  5,
  }

  // create limit store
  store := memory.NewStoreWithOptions(limiter.StoreOptions{
    MaxRetry: 5,
    CleanUpInterval: time.Duration(time.Minute*10),
  })

  // Create a new middleware with the limiter instance.
  limitRequest := mgin.NewMiddleware(limiter.New(store, rate))


  r.ForwardedByClientIP = true
  r.Use(limitRequest)
  r.Use(stdmdw.ContextWithTimeOut())
  r.Use(stdmdw.MultiUnblockCors())
  r.GET("/ping",func(ctx *gin.Context) {ctx.JSON(http.StatusOK,"pong")})

  routes.UserRoutes(r,db)
  routes.ArticleRoutes(r,db)
  routes.AuthRoutes(r,db)
  srv := &http.Server{
		Addr:    ":"+newConfig.Port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
  // Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
