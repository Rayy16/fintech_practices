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

// AuthHandler godoc
// @Summary 登录接口
// @Schemes
// @Description 登录以获取token
// @Tags user
// @Accept json
// @Param user_account body schema.AuthReq true "用户账号与加密的用户密码"
// @Produce json
// @Success 200 {object} schema.AuthResp
// @Router /login [post]
func AuthHandler(c *gin.Context) {
	log := global.Log.Sugar()
	var user schema.AuthReq
	var resp schema.AuthResp
	resp.CommResp = schema.DefaultCommResp
	err := c.ShouldBind(&user)
	if err != nil {
		log.Errorf("c.ShouldBind error: %v", err)
		resp.Msg = fmt.Sprintf("invalid params: %s", err.Error())
		resp.Code = global.REQUEST_PARAMS_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}
	log.Infof("parans: %+v user: %+v", c.Params, user)
	// ===== 改成明文传输密码 =====

	password, err := user.DecryptData, nil
	// password, err := tools.Decrypt(user.DecryptData)

	// ===========================
	if err != nil {
		log.Errorf("tools.Decrypt error: %s", err.Error())
		resp.Msg = fmt.Sprintf("decrtpt err: %s", err.Error())
		resp.Code = global.DECRYPT_DATA_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	// query account & password and check them
	md5Pw := tools.GenMD5WithSalt(string(password), tools.Salt)
	if ok, err := dao.ComparePassword(user.UserAccount, md5Pw); err != nil {
		log.Errorf("dao.ComparePassword error: %s", err.Error())
		resp.Msg = fmt.Sprintf("authorized error: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR
		c.JSON(http.StatusOK, resp)
		return
	} else if !ok {
		resp.Msg = "invalid user password"
		resp.Code = global.INVALID_PASSWORD_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	tokenString, _ := tools.GenToken(user.UserAccount)
	resp.Token = tokenString
	c.JSON(http.StatusOK, resp)
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
