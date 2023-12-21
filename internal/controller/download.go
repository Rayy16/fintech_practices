package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
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
	FtypeAudio      = fileType{"audio"}
)

// DownloadHandler godoc
// @Summary 下载文件接口
// @Schemes
// @Description 下载文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载
// @Tags download
// @Accept json
// @Param file_type path string true "下载的文件类型, 类型为枚举值：dp、resource、cover_image"
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
	var cnt int64

	switch fileTypeStr {
	case FtypeDp.String():
		cnt, _ = dao.Count(
			(&model.DigitalPersonInfo{}).TableName(),
			dao.OwnerBy(userAccount), dao.DpLinkBy(fileName), dao.StatusBy(dao.StatusSuccess),
		)
	case FtypeCoverImage.String():
		cnt, _ = dao.Count(
			(&model.DigitalPersonInfo{}).TableName(),
			dao.OwnerBy(userAccount), dao.CoverImageLinkBy(fileName), dao.StatusBy(dao.StatusSuccess),
		)
	case FtypeResource.String():
		cnt, _ = dao.Count(
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

	if cnt == 0 {
		log.Infof("file %s not belong to %s Or not existed", fileName, userAccount)
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("file %s not belong to %s Or not existed", fileName, userAccount),
			"code": global.AUTHORIZATION_ERROR,
		})
		return
	}

	var err error
	rootDir := global.RootDirMap[fileTypeStr]
	filepath := path.Join(rootDir, fileName)
	_, err = os.Stat(filepath)
	if err != nil {
		log.Errorf("os.Stat(%s) error: %s", fileName, err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("can not find file %s err: %s", fileName, err.Error()),
			"code": global.INVALID_FILE_ERROR,
		})
		return
	}
	c.File(filepath)
}
