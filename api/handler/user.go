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

type UserHandler struct {
	Db *database.DB
}

func NewUserHandler(db *database.DB) (UserHandler, error) {
	return UserHandler{Db: db}, nil
}

func (users *UserHandler) CreateUserHandler(c *gin.Context) {
	// Parse the request body to get the user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// prepare data befor insert to database
	newUser := models.User{
		Username:  user.Username,
		Password:  user.Password,
	}

	// insert data to database
	err := users.Db.CreateUser(&newUser)
	if err != nil {
		errMsg := utils.InternalServerError("can't insert data to database or username already exit")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg)
		return
	}
	// Return a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (users *UserHandler) UpdateUserHandler(c *gin.Context) {
	// get id from param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errMsg := utils.BadRequestError("id should be number only")
		c.JSON(errMsg.Status, errMsg.Message)
	}

	// Parse the request body to get the user data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user data
	oldUser, err := users.Db.GetUser(strconv.Itoa(id))
	if err != nil {
		errMsg := utils.BadRequestError("does't exit")
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}

	// prepare data befor insert to database
	newUser := models.User{
		Username:  user.Username,
		Password:  user.Password,
	}

	if user.Username == "" {
		newUser.Username = oldUser.Username
	}
	if user.Password == "" {
		newUser.Password = oldUser.Password
	}

	// insert data to database
	err = users.Db.UpdateUser(&newUser)
	if err != nil {
		errMsg := utils.InternalServerError("can't update data to database")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg)
		return
	}
	// Return a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (users *UserHandler) GetUserHandler(c *gin.Context) {
	// Get the user ID from the request parameters
	userID := c.Param("id")
	// Retrieve the user from the database or any other data source

	user, err := users.Db.GetUser(userID)
	if err != nil {
		errMsg := utils.InternalServerError("can't query data from database")
		log.Println(errMsg.Message)
		c.JSON(errMsg.Status, errMsg.Message)
	}

	// Return the user as JSON response
	c.JSON(http.StatusOK, user)
}

func (users *UserHandler) GetUsersHandler(c *gin.Context) {
	user, err := users.Db.GetUsers()
	if err != nil {
		errMsg := utils.InternalServerError("can't query data in database")
		c.JSON(errMsg.Status, errMsg.Message)
	}
	// Retrieve the user from the database or any other data source
	// For demonstration purposes, let's assume we have a user with ID 1

	// Return the user as JSON response
	c.JSON(http.StatusOK, user)
}

func (users *UserHandler) DeleteUserHandler(c *gin.Context) {
	// Get the user ID from the request parameters
	userID := c.Param("id")

	// call database
	err := users.Db.DeleteUser(userID)
	if err != nil {
		errMsg := utils.InternalServerError("something wrong in server")
		log.Println(err)
		c.JSON(errMsg.Status, errMsg.Message)
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

