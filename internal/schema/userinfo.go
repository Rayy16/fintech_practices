package schema

type AuthReq struct {
	UserAccount string `json:"user_account" binding:"required"`
	DecryptData string `json:"decrypt_data" binding:"required"`
}

type RegisterReq struct {
	UserAccount string `json:"user_account" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	DecryptData string `json:"decrypt_data" binding:"required"`
}
