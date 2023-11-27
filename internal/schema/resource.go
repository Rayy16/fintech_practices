package schema

import "time"

type GetResourceReq struct {
	PageParams
	IsPublic bool `json:"is_public" form:"is_public"`
}

type GetResourceResp struct {
	CommResp
	Data []ResourceEntity `json:"data"`
}

type ResourceEntity struct {
	ResouceId        string     `json:"resouce_id"`
	ResourceDescribe string     `json:"resource_describe"`
	ResourceLink     string     `json:"resource_link"`
	DpCoverImageLink string     `json:"cover_image_link"`
	CreateTime       *time.Time `json:"create_time"`
	UpdateTime       *time.Time `json:"update_time"`
}
