package controller

import (
	"errors"
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"fintechpractices/internal/schema"
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type fileType struct {
	t string
}

func (f fileType) String() string {
	return f.t
}

var (
	FtypeDp         = fileType{"dp"}
	FtypeCoverImage = fileType{"cover_image"}
	FtypeResource   = fileType{"resource"}
)

// /dp/demo.mp4; /cover_image/demo.png; /resource/demo.wav; /resource/demo.png
func DownloadHandler(c *gin.Context) {
	log := global.Log.Sugar()
	fileTypeStr := c.Param("file_type")
	fileName := c.Param("file_name")
	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)

	// query file belong to user account
	var err error

	switch fileTypeStr {
	case FtypeDp.String():
		_, err = dao.Count(
			(&model.DigitalPersonInfo{}).TableName(),
			dao.OwnerBy(userAccount), dao.DpLinkBy(fileName), dao.StatusBy(dao.StatusSuccess),
		)
	case FtypeCoverImage.String():
		_, err = dao.Count(
			(&model.DigitalPersonInfo{}).TableName(),
			dao.OwnerBy(userAccount), dao.CoverImageLinkBy(fileName), dao.StatusBy(dao.StatusSuccess),
		)
	case FtypeResource.String():
		_, err = dao.Count(
			(&model.MetadataMarket{}).TableName(),
			dao.OwnerBy(userAccount), dao.ResourceLinkBy(fileName),
		)
	default:
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("invalid file type: %s", fileTypeStr),
			"code": global.INVALID_FILE_ERROR,
		})
		return
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Infof("file %s not belong to %s Or not existed", fileName, userAccount)
			c.JSON(http.StatusOK, gin.H{
				"msg":  fmt.Sprintf("file %s not belong to %s Or not existed", fileName, userAccount),
				"code": global.AUTHORIZATION_ERROR,
			})
		} else {
			log.Errorf("check file %s existed by quering db error: %s", fileName, err.Error())
			c.JSON(http.StatusOK, gin.H{
				"msg":  fmt.Sprintf("query db error: %s", err.Error()),
				"code": global.DAO_LAYER_ERROR,
			})
		}
		return
	}

	rootDir := global.RootDirMap[fileTypeStr]
	filepath := path.Join(rootDir, fileName)
	c.JSON(http.StatusOK, schema.DefaultCommResp)
	c.File(filepath)
}
