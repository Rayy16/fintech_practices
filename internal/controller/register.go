package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"fintechpractices/internal/schema"
	"fintechpractices/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary 注册用户接口
// @Description 注册用户
// @schema http
// @Tags user
// @Accept json
// @Produce json
// @Param user_name body schema.RegisterReq true "用户名称、用户账户与加密的用户密码"
// @Success 200 {object} schema.CommResp
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	log := global.Log.Sugar()
	var req schema.RegisterReq
	resp := schema.DefaultCommResp

	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("c.ShouldBind err: %s", err.Error())
		resp.Msg = err.Error()
		resp.Code = global.REQUEST_PARAMS_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	if exited, err := dao.IsAccountExisted(req.UserAccount); err != nil {
		log.Errorf("dao.IsAccountExisted err: %s", err.Error())
		resp.Msg = err.Error()
		resp.Code = global.REQUEST_PARAMS_ERROR
		c.JSON(http.StatusOK, resp)
		return
	} else if exited {
		log.Infof("user account %s already exists", req.UserAccount)
		resp.Msg = fmt.Sprintf("user account %s already exists", req.UserAccount)
		resp.Code = global.REQUEST_PARAMS_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	var user model.UserInfo
	// ===== 改成明文传输密码 =====
	var password string
	var err error
	password, err = req.DecryptData, nil
	// password, err := tools.Decrypt(req.DecryptData)

	// ===========================
	if err != nil {
		log.Errorf("tools.Decrypt err: %s", err.Error())
		resp.Msg = fmt.Sprintf("decrypt data err: %s", err.Error())
		resp.Code = global.DECRYPT_DATA_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	md5Password := tools.GenMD5WithSalt(string(password), tools.Salt)
	user.UniAccount = req.UserAccount
	user.UserName = req.UserName
	user.PassWord = md5Password

	if err := dao.CreateUser(&user); err != nil {
		log.Errorf("dao.CreateUser err: %s", err.Error())
		resp.Msg = fmt.Sprintf("create user %s err: %s", req.UserAccount, err.Error())
		resp.Code = global.DAO_LAYER_ERROR
	}
	c.JSON(http.StatusOK, resp)
}
