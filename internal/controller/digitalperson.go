package controller

import (
	"fintechpractices/global"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"fintechpractices/internal/schema"
	"fintechpractices/internal/task"
	"fintechpractices/internal/task/types"
	"fintechpractices/tools"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

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

// GetDpHandler godoc
// @Summary 查询数字人接口
// @Schemes
// @Description 查询用户所拥有的数字人信息
// @Tags digital person
// @Accept json
// @Produce json
// @Param page_no query int false "分页查询页数，默认为1"
// @Param page_size query int false "分页查询页大小，默认为10"
// @Param order_field query string false "查询返回的排序字段，默认为创建时间"
// @Param method query string false "排序方式，默认为倒序"
// @Param Authorization header string true "token"
// @Success 200 {object} schema.GetDpResp
// @Router /dp [get]
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
		dao.OwnerBy(userAccount), dao.OrderBy(field+method), dao.PageBy(req.PageNo, req.PageSize),
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
			DpStatus:       dps[i].DpStatus,
			CreateTime:     dps[i].CreateTime,
			UpdateTime:     dps[i].UpdateTime,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteDpHandler godoc
//
// @Summary 删除数字人接口
// @Schemes
// @Description 删除用户所拥有的数字人信息
// @Tags digital person
// @Accept json
// @Produce json
// @Param dp_link path string true "需删除的数字人id"
// @Param Authorization header string true "token"
// @Success 200 {object} schema.CommResp
// @Router /dp/{dp_link} [delete]
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

func CreateDpHandler(c *gin.Context) {
	log := global.Log.Sugar()
	var req schema.CreateDpReq
	var resp = schema.DefaultCommResp

	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("c.ShouldBind err: %s", err.Error())
		resp.Code = global.REQUEST_PARAMS_ERROR
		resp.Msg = fmt.Sprintf("binding body params error: %s", err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}

	// 判断是否需要调用模型生成Audio
	args := make([]types.TaskArgs, 0)
	var dpLink string
	var coverImageLink string
	if req.AudioLink == "" {
		req.AudioLink = tools.GenMD5(req.Content+req.ToneLink) + ".wav"
		args = append(args, types.AudioArgs{
			TextInput: req.Content,
			ToneInput: filepath.Join(global.RootDirMap[FtypeResource.String()], req.ImageLink),
			OutputDir: global.RootDirMap[FtypeAudio.String()],
			FileName:  req.AudioLink,
		})
	}
	raw, _ := c.Get("user_account")
	userAccount, _ := raw.(string)
	dpLink = tools.GenMD5(userAccount+req.ImageLink+req.AudioLink) + ".mp4"
	coverImageLink = tools.GenMD5(dpLink+"_cover_image") + ".png"
	args = append(args, types.DpArgs{
		AudioInput: filepath.Join(global.RootDirMap[FtypeAudio.String()], req.AudioLink),
		ImageInput: filepath.Join(global.RootDirMap[FtypeResource.String()], req.ImageLink),
		OutputDir:  global.RootDirMap[FtypeDp.String()],
		FileName:   dpLink,
	})
	dp := &model.DigitalPersonInfo{
		DpId:           dpLink,
		DpName:         req.DpName,
		DpStatus:       dao.StatusCreatable.Int(),
		OwnerId:        userAccount,
		Published:      true,
		Audited:        true,
		Content:        req.Content,
		CoverImageLink: coverImageLink,
		DpLink:         dpLink,
	}
	err := dao.CreateDigitalPerson(dp)
	if err != nil {
		resp.Code = global.DAO_LAYER_ERROR
		resp.Msg = fmt.Sprintf("create dp err: %s", err.Error())
		log.Errorf("dao.CreateDigitalPerson err: %s", err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}

	global.TaskMgr.RegisterTask(dpLink, args...)
	go func() {
		for {
			raw, _ := global.TaskMgr.QueryTask(dpLink)
			info, _ := raw.(task.TaskInfo)
			switch info.Status {
			case dao.StatusFailed.Int():
				dao.UpdateDPStatusByLink(dpLink, dao.StatusFailed)
				log.Errorf("dp <%s> create failed: %s", dpLink, info.Msg)
				return
			case dao.StatusSuccess.Int():
				dpPath := filepath.Join(global.RootDirMap[FtypeDp.String()], dpLink)
				coverImagePath := filepath.Join(global.RootDirMap[FtypeCoverImage.String()], coverImageLink)
				err := tools.ExtractVedioToImage(dpPath, coverImagePath)
				if err != nil {
					log.Errorf("failed to tract dp <%s> cover image: %s", dpLink, err.Error())
				}
				dao.UpdateDPStatusByLink(dpLink, dao.StatusSuccess)
				log.Infof("dp <%s> create success", dpLink)
			default:
				time.Sleep(time.Second)
			}
		}
	}()
	c.JSON(http.StatusOK, resp)
}
