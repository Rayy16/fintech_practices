package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/schema"
	"fintechpractices/tools"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	log := global.Log.Sugar()
	fileTypeStr := c.Param("file_type")
	fileName := c.Param("file_name")
	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)
	var err error
	var resp = schema.DefaultCommResp

	// 判断是否为素材的所有者
	rs, cnt, err := dao.GetResourceBy(dao.OwnerBy(userAccount), dao.ResourceLinkBy(fileName))
	if err != nil {
		log.Errorf("dao.GetResourceBy(%s, %s) err: %s", userAccount, fileName, err.Error())
		resp.Code = global.DAO_LAYER_ERROR
		resp.Msg = fmt.Sprintf("get resource by %s err: %s", userAccount, err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}
	if cnt == 0 {
		log.Infof("dao.GetResourceBy(%s, %s) not found", userAccount, fileName)
		resp.Code = global.AUTHORIZATION_ERROR
		resp.Msg = "authorization err: record not found"
		c.JSON(http.StatusOK, resp)
		return
	}
	resource := rs[0]

	file, err := c.FormFile("file")
	if err != nil {
		resp.Msg = fmt.Sprintf("extract file from form err: %s", err.Error())
		resp.Code = global.INVALID_FILE_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}
	filePath := filepath.Join(global.RootDirMap[fileTypeStr], fileName)
	switch fileTypeStr {
	case FtypeAudio.String():
		c.SaveUploadedFile(file, filePath)
		log.Infof("upload resource <%s> file: %s", resource.ResourceType, fileName)
	case FtypeResource.String():
		c.SaveUploadedFile(file, filePath)
		log.Infof("upload resource <%s> file: %s", resource.ResourceType, fileName)
		// 如果是形象素材，则可能需要截取封面
		if resource.ResourceType == dao.TypeImage.String() {
			// 如果形象素材不是视频形象，那么跳过即可
			if !strings.Contains(resource.ResourceLink, ".mp4") {
				break
			}
			// 如果形象素材是视频形象，那么需要截取封面
			coverImagePath := filepath.Join(global.RootDirMap[FtypeCoverImage.String()], resource.CoverImageLink)
			err = tools.ExtractVedioToImage(filePath, coverImagePath)
			if err != nil {
				resp.Code = global.INVALID_FILE_ERROR
				resp.Msg = fmt.Sprintf("tract vedio to image err: %s", err.Error())
				c.JSON(http.StatusOK, resp)
			}
		}
		err = dao.UpdateResourceByLink(fileName, map[string]interface{}{"is_ready": true})
	default:
		resp.Code = global.INVALID_FILE_ERROR
		resp.Msg = fmt.Sprintf("invalid file type: %s", fileTypeStr)
	}
	if err != nil {
		resp.Code = global.DAO_LAYER_ERROR
		resp.Msg = fmt.Sprintf("update resource status failed: %s", err.Error())
		os.Remove(filePath)
	}
	c.JSON(http.StatusOK, resp)
}
