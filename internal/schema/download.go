package schema

type DownloadReq struct {
	FileLink string `json:"file_link" form:"file_link"`
	FileType string `json:"file_type" form:"file_type"`
}
