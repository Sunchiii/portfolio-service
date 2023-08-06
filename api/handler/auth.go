package handler

import (
	"net/http"
	"strconv"
  "log"

	"github.com/gin-gonic/gin"
	"github.com/sunchiii/portfolio-service/api/middleware"
	"github.com/sunchiii/portfolio-service/api/models"
	"github.com/sunchiii/portfolio-service/pkg/database"
	"github.com/sunchiii/portfolio-service/pkg/utils"
)

type AuthHandler struct{
  DB *database.DB
}

func NewAuthHandler(db *database.DB) (AuthHandler,error){
  return AuthHandler{
    DB: db,
  },nil
}

func (authdb *AuthHandler) Login(c *gin.Context){
  // prepare data when user request
  var user models.Login
  
  // get data json from user
  if err := c.ShouldBindJSON(&user); err != nil{
    errMsg := utils.BadRequestError("please bind username and password")
    c.JSON(errMsg.Status,errMsg.Message)
    return
  }
  
  // find user from database 
  userdb,err := authdb.DB.GetUserByUsername(user.Username, user.Password)
  if err != nil{
    errMsg := utils.BadRequestError("username does't exit or incorrect password")
    log.Println(err)
    c.JSON(errMsg.Status,errMsg.Message)
    return
  }

  token,err := middleware.GenerateToken(strconv.Itoa(int(userdb.ID)))
  if err != nil{
    errMsg := utils.InternalServerError("something wrong in server")
    c.JSON(errMsg.Status,errMsg.Message)
    return
  }

  newUser := models.User{
    ID: userdb.ID,
    Username: userdb.Username,
    CreatedAt: userdb.CreatedAt,
    Password: "asefksaefmnfknjasefjnsae,nfms;keanfsaenfskenfase",
  }

  c.JSON(http.StatusOK,gin.H{"user_data":newUser,"token":token})

}
