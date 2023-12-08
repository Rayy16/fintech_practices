package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/schema"
	"fintechpractices/tools"
)

func AuthHandler(c *gin.Context) {
	log := global.Log.Sugar()
	var user schema.AuthReq
	err := c.ShouldBind(&user)
	if err != nil {
		log.Errorf("c.ShouldBind error: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("invalid params: %s", err.Error()),
			"code": global.REQUEST_PARAMS_ERROR,
		})
	}
	password, err := tools.Decrypt(user.DecryptData)
	if err != nil {
		log.Errorf("tools.Decrypt error: %s", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("decrtpt err: %s", err.Error()),
			"code": global.DECRYPT_DATA_ERROR,
		})
		return
	}

	// query account & password and check them
	md5Pw := tools.GenMD5WithSalt(string(password), tools.Salt)
	if ok, err := dao.ComparePassword(user.UserAccount, md5Pw); err != nil {
		log.Errorf("dao.ComparePassword error: %s", err.Error)
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("authorized error: %s", err.Error()),
			"code": global.DAO_LAYER_ERROR,
		})
		return
	} else if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "invalid user password",
			"code": global.INVALID_PASSWORD_ERROR,
		})
		return
	}

	tokenString, _ := tools.GenToken(user.UserAccount)
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"code":  global.SUCCESS,
		"token": tokenString,
	})
}

func GetAuthMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "request header without authorization",
				"code": global.AUTHORIZATION_ERROR,
			})
			c.Abort()
			return
		}

		raws := strings.SplitN(authHeader, " ", 2)
		if len(raws) != 2 || raws[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "request header with invalid auth type",
				"code": global.AUTHORIZATION_ERROR,
			})
			c.Abort()
			return
		}

		token := raws[1]
		claims, err := tools.ParseToken(token)
		if err != nil || claims.ExpiresAt.Before(time.Now()) || claims.NotBefore.After(time.Now()) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  fmt.Sprintf("parse token error: %s", err.Error()),
				"code": global.AUTHORIZATION_ERROR,
			})
			c.Abort()
			return
		}
		c.Set("user_account", claims.UserAccount)
		c.Next()
	}
}
