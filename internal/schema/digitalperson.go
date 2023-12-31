package schema

import "time"

type DpEntity struct {
	DpName         string     `json:"dp_name"`
	Owner          string     `json:"owner"`
	DpLink         string     `json:"dp_link"`
	CoverImageLink string     `json:"cover_image_link"`
	HotScore       int        `json:"hot_score"`
	DpStatus       int        `json:"dp_status"`
	CreateTime     *time.Time `json:"create_time"`
	UpdateTime     *time.Time `json:"update_time"`
}

type GetDpReq struct {
	PageParams
	OrderField string `json:"order_field" form:"order_field"`
	Method     string `json:"method" form:"order_field"`
}

type GetDpResp struct {
	CommResp
	Data []DpEntity `json:"data"`
}

type CreateDpReq struct {
	DpName    string `json:"dp_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
	AudioLink string `json:"audio_link" binding:"required"`
	ToneLink  string `json:"tone_link" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
