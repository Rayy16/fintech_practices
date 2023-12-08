package controller

import (
	"errors"
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
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

// DownloadHandler godoc
// @Summary 下载文件接口
// @Schemes
// @Description 下载文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载
// @Tags download
// @Accept json
// @Param file_type path string true "下载的文件类型"
// @Param file_name path string true "下载的文件名称"
// @Param Authorization header string true "token"
// @Produce */*
// @Success 200 {file} file
// @Router /{file_type}/{file_name} [get]
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
	c.File(filepath)
}
