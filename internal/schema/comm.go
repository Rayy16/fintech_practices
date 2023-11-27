package schema

import "fintechpractices/global"

var (
	DefaultCommResp = CommResp{
		Msg:  "success",
		Code: global.SUCCESS,
	}
)

type CommResp struct {
	Msg  string `json:"msg" form:"msg"`
	Code int    `json:"code" form:"code"`
}

type PageParams struct {
	PageNo   int `json:"page_no" form:"page_no"`
	PageSize int `json:"page_size" form:"page_size"`
}
