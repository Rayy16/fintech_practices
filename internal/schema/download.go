package schema

type DownloadReq struct {
	FileLink string `json:"file_link" form:"file_link" binding:"required"`
	FileType string `json:"file_type" form:"file_type" binding:"required"`
}
