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

// UploadHandler godoc
// @Summary 上传文件接口
// @Schemes
// @Description 上传文件的统一接口，数字人、封面图片、素材库素材均通过本接口下载
// @Tags download
// @Accept json
// @Param file_type path string true "上传的文件类型，类型为枚举值：audio、resource"
// @Param file_name path string true "上传的文件名称，需要带上相应后缀，例如audio为.wav, resource为 .png 或 .mp4"
// @Param file formData file true "上传的文件"
// @Param Authorization header string true "token"
// @Produce */*
// @Success 200 {object} schema.CommResp
// @Router /upload/{file_type}/{file_name} [post]
func UploadHandler(c *gin.Context) {
	log := global.Log.Sugar()
	fileTypeStr := c.Param("file_type")
	fileName := c.Param("file_name")
	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)
	var err error
	var resp = schema.DefaultCommResp

	filePath := filepath.Join(global.RootDirMap[fileTypeStr], fileName)
	file, err := c.FormFile("file")
	if err != nil {
		resp.Msg = fmt.Sprintf("extract file from form err: %s", err.Error())
		resp.Code = global.INVALID_FILE_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	switch fileTypeStr {
	case FtypeAudio.String():
		c.SaveUploadedFile(file, filePath)
		log.Infof("upload audio file: %s", fileName)
	case FtypeResource.String():
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

		c.SaveUploadedFile(file, filePath)
		log.Infof("upload resource <%s> file: %s", resource.ResourceType, fileName)
		// 如果是视频形象素材，则需要截取封面
		// 图片形象素材 以及 音色素材则跳过
		if resource.ResourceType == dao.TypeImage.String() && strings.Contains(resource.ResourceLink, ".mp4") {
			coverImagePath := filepath.Join(global.RootDirMap[FtypeCoverImage.String()], resource.CoverImageLink)
			err = tools.ExtractVedioToImage(filePath, coverImagePath)
			if err != nil {
				resp.Code = global.INVALID_FILE_ERROR
				resp.Msg = fmt.Sprintf("tract vedio to image err: %s", err.Error())
				c.JSON(http.StatusOK, resp)
			}
		}
		// 最后都要经过 素材状态 的更新方法
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
