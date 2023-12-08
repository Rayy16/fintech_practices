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

var (
	OrderHotScore   = orderField{"hot_score"}
	OrderCreateTime = orderField{"create_time"}
	OrderUpdateTime = orderField{"update_time"}
)

type orderField struct {
	t string
}

func (o orderField) String() string {
	return o.t
}

// /dp/page_no=?&page_size=?order_field=?method=?
func GetDpHandler(c *gin.Context) {
	log := global.Log.Sugar()

	var req = schema.GetDpReq{
		PageParams: schema.PageParams{PageNo: 1, PageSize: 10},
		OrderField: OrderCreateTime.String(),
	}
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("c.ShouldBind error: %s", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("binding query param err: %s", err.Error()),
			"code": global.REQUEST_PARAMS_ERROR,
		})
		return
	}

	var field string
	switch req.OrderField {
	case OrderCreateTime.String(), OrderUpdateTime.String(), OrderHotScore.String():
		field = req.OrderField
	case "":
		field = OrderCreateTime.String()
	default:
		c.JSON(http.StatusOK, gin.H{
			"msg":  fmt.Sprintf("invliad query param: order_field=%s", req.OrderField),
			"code": global.REQUEST_PARAMS_ERROR,
		})
		return
	}
	var method = " desc"
	if strings.ToLower(req.Method) == "asc" {
		method = "asc"
	}

	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)

	dps, cnt, err := dao.GetDigitalPersonsBy(
		dao.OwnerBy(userAccount), dao.AuditedBy(true), dao.OrderBy(field+method), dao.PageBy(req.PageNo, req.PageSize),
	)

	var resp schema.GetDpResp
	resp.CommResp = schema.DefaultCommResp

	if err != nil {
		log.Errorf("dao.GetDigitalPersonBy err: %s", err.Error())
		resp.Msg = fmt.Sprintf("get digital person err: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR
		c.JSON(http.StatusOK, resp)
		return
	}

	if cnt == 0 {
		resp.Msg = "record not found"
	}

	resp.Data = make([]schema.DpEntity, 0, len(dps))
	for i := range dps {
		resp.Data = append(resp.Data, schema.DpEntity{
			DpName:         dps[i].DpName,
			Owner:          dps[i].OwnerId,
			DpLink:         dps[i].DpLink,
			CoverImageLink: dps[i].CoverImageLink,
			HotScore:       dps[i].HotScore,
			CreateTime:     dps[i].CreateTime,
			UpdateTime:     dps[i].UpdateTime,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// /dp/:dp_link
func DeleteDpHandler(c *gin.Context) {
	log := global.Log.Sugar()
	dpLink := c.Param("dp_link")
	err := dao.DeleteDigitalPersonByLink(dpLink)

	var resp = schema.DefaultCommResp
	if err != nil {
		log.Errorf("dao.DeleteDigitalPersonByLink err: %s", err.Error())
		resp.Msg = fmt.Sprintf("delete digital person err: %s", err.Error())
		resp.Code = global.DAO_LAYER_ERROR
	}

	c.JSON(http.StatusOK, resp)
}
