package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fintechpractices/global"
	"fintechpractices/internal/schema"
	"fintechpractices/tools"
)

// PubKeyHandler godoc
// @Summary 获取公钥接口
// @schema http
// @Description 获取rsa公钥
// @Tags authorization
// @Accept json
// @Produce json
// @Success 200 {object} schema.PubKeyResp
// @Router /pubkey [get]
func PubKeyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, schema.PubKeyResp{
		CommResp: schema.CommResp{
			Msg:  "success",
			Code: global.SUCCESS,
		},
		Data: tools.GetPubKey(),
	})
}
