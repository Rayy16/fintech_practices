package controller

import (
	"errors"
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/schema"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// /resource/:resource_type/page_no=?&page_size=?&is_public=?
func GetResourceHandler(c *gin.Context) {
	log := global.Log.Sugar()

	resourceTypeStr := strings.ToLower(c.Param("resource_type"))
	resourceType, err := dao.NewResourceType(resourceTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("get resource type err: %s", err.Error()),
			"code": global.REQUEST_PARAMS_ERROR,
		})
		return
	}

	var req = schema.GetResourceReq{
		PageParams: schema.PageParams{
			PageNo:   1,
			PageSize: 10,
		},
		IsPublic: false,
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("binding query param err: %s", err.Error()),
			"code": global.REQUEST_PARAMS_ERROR,
		})
		return
	}

	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)
	if req.IsPublic {
		userAccount = global.PublicAccount
	}
	rs, err := dao.GetResourceBy(
		dao.OwnerBy(userAccount), dao.TypeBy(resourceType), dao.PageBy(req.PageNo, req.PageSize),
	)

	var resp schema.GetResourceResp
	resp.CommResp = schema.DefaultCommResp

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Msg = "record not found"
		} else {
			log.Errorf("dao.GetResourceBy err: %s", err.Error())
			resp.Msg = fmt.Sprintf("get resource err: %s", err.Error())
			resp.Code = global.DAO_LAYER_ERROR
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = make([]schema.ResourceEntity, 0, len(rs))
	for i := range rs {
		var cImageLink string = rs[i].ResourceLink
		if rs[i].CoverImageLink != "" {
			cImageLink = rs[i].CoverImageLink
		}

		resp.Data = append(resp.Data, schema.ResourceEntity{
			ResouceId:        rs[i].ResourceId,
			ResourceDescribe: rs[i].ResourceDescribe,
			ResourceLink:     rs[i].ResourceLink,
			DpCoverImageLink: cImageLink,
			CreateTime:       rs[i].CreateTime,
			UpdateTime:       rs[i].UpdateTime,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// /resource/:resource_link
func DeleteResourceHandler(c *gin.Context) {
	log := global.Log.Sugar()
	rsLink := c.Param("resource_link")
	err := dao.DeleteResourceByLink(rsLink)

	var resp = schema.DefaultCommResp
	if err != nil {
		log.Errorf("dao.DeleteResourceByLink err: %s", err.Error())
		resp.Msg = fmt.Sprintf("delete resource err: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR
	}
	c.JSON(http.StatusOK, resp)
}
