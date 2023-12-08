package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/schema"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetResourceHandler godoc
// @Summary 查询素材库接口
// @Schemes
// @Description 查询用户所拥有的 or 公共的素材库素材信息
// @Tags resource lib
// @Accept json
// @Produce json
// @Param resource_type path string true "素材类型，tone or image"
// @Param page_no query int false "分页查询页数，默认为1"
// @Param page_size query int false "分页查询页大小，默认为10"
// @Param is_public query boolean false "是否查询公共素材，默认为否"
// @Param Authorization header string true "token"
// @Success 200 {object} schema.GetResourceResp
// @Router /resource/{resource_type} [get]
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
		log.Errorf("c.ShouldBind error: %s", err.Error())
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
	rs, cnt, err := dao.GetResourceBy(
		dao.OwnerBy(userAccount), dao.TypeBy(resourceType), dao.PageBy(req.PageNo, req.PageSize),
	)

	var resp schema.GetResourceResp
	resp.CommResp = schema.DefaultCommResp

	if err != nil {
		log.Errorf("dao.GetResourceBy err: %s", err.Error())
		resp.Msg = fmt.Sprintf("get resource err: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}
	if cnt == 0 {
		resp.Msg = "record not found"
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

// DeleteResourceHandler godoc
// @Summary 删除素材库素材接口
// @Schemes
// @Description 删除用户所拥有的素材库素材
// @Tags resource lib
// @Accept json
// @Produce json
// @Param resource_link path string true "素材连接，tone or image"
// @Param Authorization header string true "token"
// @Success 200 {object} schema.CommResp
// @Router /resource/{resource_link} [delete]
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
