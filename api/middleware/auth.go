package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	paseto "github.com/o1egl/paseto"
	"github.com/sunchiii/portfolio-service/config"
	"github.com/sunchiii/portfolio-service/pkg/utils"
)

func AuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
    baererToken := c.GetHeader("Authorization")
    tokenString := baererToken[7:]
    if tokenString == ""{
      errMsg := utils.UnauthorizedError("token not found")
      c.JSON(errMsg.Status,errMsg.Message)
      c.Abort()
      return
    }
		// get signature from config
		pasetoConfig, err := config.GetPasetoConfig()
		if err != nil {
      errMsg := utils.InternalServerError("something wrong in server side")
      c.JSON(errMsg.Status,errMsg.Message)
      c.Abort()
      return
		}
		var newJson paseto.JSONToken
		var newFooter string
		err = paseto.NewV1().Decrypt(tokenString,[]byte(pasetoConfig.SignatureKey),&newJson,&newFooter)
    if err != nil{
      errMsg := utils.UnauthorizedError("token invalid please login")
      c.JSON(errMsg.Status,errMsg.Message)
      c.Abort()
      return
    }

		c.Next()
	}
}

func GenerateToken(userId string) (string, error) {
	// get time now
	now := time.Now()
	// get signature from config
	pasetoConfig, err := config.GetPasetoConfig()
	if err != nil {
		return "", err
	}

	// convert to byte 32
	symetric := []byte(pasetoConfig.SignatureKey)

	// make token payload
	jsonToken := paseto.JSONToken{
		IssuedAt:   now,
		Expiration: now.Add(time.Duration(pasetoConfig.ExpHour) * time.Hour),
		NotBefore:  now,
	}

	// // custome claim
	jsonToken.Set("data", pasetoConfig.SignatureKey)
	flooter := "signature footer"

	// enscript data
	token, err := paseto.NewV1().Encrypt(symetric, jsonToken, flooter)
	if err != nil {
		return "", err
	}

	return token, nil

}
