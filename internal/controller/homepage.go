package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// /hotvedio/pageNo=?&pageSize=?
// HotVedioHandler godoc
// @Summary 首页视频接口
// @Schemes
// @Description 获取首页数字人视频
// @Tags homepage
// @Accept json
// @Param pageNo query string false "分页查询，页码"
// @Param pageSize query string false "分页查询，页大小"
// @Produce */*
// @Success 200 {object} schema.GetDpResp
// @Router /hotvedio [get]
func HotVedioHandler(c *gin.Context) {
	log := global.Log.Sugar()
	var req = schema.PageParams{PageNo: 1, PageSize: 5}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("binding query param err: %s", err.Error()),
			"code": global.REQUEST_PARAMS_ERROR,
		})
		return
	}

	dps, cnt, err := dao.GetDigitalPersonsBy(dao.StatusBy(dao.StatusSuccess), dao.PublishedBy(true), dao.AuditedBy(true),
		dao.OrderBy("hot_score desc"), dao.OrderBy("create_time"), dao.PageBy(req.PageNo, req.PageSize))

	var resp schema.GetDpResp
	resp.CommResp = schema.DefaultCommResp

	if err != nil {
		log.Errorf("dao.GetDigitalPersonBy err: %s", err.Error())
		resp.Msg = fmt.Sprintf("get hot vedio err: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR

		c.JSON(http.StatusOK, resp)
		return
	}
	if cnt == 0 {
		resp.Msg = "record not found"
	}

	resp.Data = make([]schema.DpEntity, 0, len(dps))
	for i := range dps {
		var owner string = dps[i].OwnerId
		if uInfo, err := dao.GetUserInfo(dps[i].OwnerId); err == nil {
			owner = uInfo.UserName
		} else {
			log.Errorf("dao.GetUserInfo err: %s", err.Error())
		}

		resp.Data = append(resp.Data, schema.DpEntity{
			DpName:         dps[i].DpName,
			Owner:          owner,
			DpLink:         dps[i].DpLink,
			CoverImageLink: dps[i].CoverImageLink,
			HotScore:       dps[i].HotScore,
			CreateTime:     dps[i].CreateTime,
			UpdateTime:     dps[i].UpdateTime,
		})
	}
	c.JSON(http.StatusOK, resp)
}
