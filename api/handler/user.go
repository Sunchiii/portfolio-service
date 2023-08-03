package handler

import (
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/models"
	"github.com/sunchiii/portfolio-service/pkg/database"
	"github.com/sunchiii/portfolio-service/pkg/utils"
)

type UserHandler struct{
  Db *database.DB
}

func NewUserHandler(db *database.DB)(UserHandler,error){
  return UserHandler{Db: db},nil
}

func createUserHandler(c *gin.Context) {
	// Parse the request body to get the user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the user to the database or any other data source
  fmt.Println(user)
	// ...

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func getUserHandler(c *gin.Context) {
	// Get the user ID from the request parameters
	// userID := c.Param("id")

	// Retrieve the user from the database or any other data source
	// For demonstration purposes, let's assume we have a user with ID 1
	user := models.User{
		ID:        1,
		Username:  "john",
		Password:  "password",
		CreatedAt: time.Now(),
	}

	// Return the user as JSON response
	c.JSON(http.StatusOK, user)
}
func (users *UserHandler) GetUsersHandler(c *gin.Context) {
  user,err := users.Db.GetUsers()
  if err != nil{
    errMsg := utils.InternalServerError("can't query data in database")
    c.JSON(errMsg.Status,errMsg.Message)
  }
	// Retrieve the user from the database or any other data source
	// For demonstration purposes, let's assume we have a user with ID 1

	// Return the user as JSON response
	c.JSON(http.StatusOK, user)
}


func deleteUserHandler(c *gin.Context) {
	// Get the user ID from the request parameters
	// userID := c.Param("id")

	// Delete the user from the database or any other data source
	// ...

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
