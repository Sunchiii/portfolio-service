package handler

import (
	"database/sql"
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/models"
)

type UserHandler struct{
  Db *sql.DB
}

func NewUserHandler(db *sql.DB)(UserHandler,error){
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

	// Retrieve the user from the database or any other data source
	// For demonstration purposes, let's assume we have a user with ID 1
  user := []models.User{
    {
	 	  ID:        1,
	 	  Username:  "john",
	 	  Password:  "password",
	 	  CreatedAt: time.Now(),
    },
  }
  

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
