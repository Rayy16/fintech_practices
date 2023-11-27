package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fintechpractices/tools"
)

func PubKeyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": tools.GetPubKey(),
	})
}
