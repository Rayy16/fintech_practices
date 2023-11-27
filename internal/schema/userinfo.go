package schema

type AuthReq struct {
	UserAccount string `json:"user_account"`
	DecryptData string `json:"decrypt_data"`
}

type RegisterReq struct {
	UserAccount string `json:"user_account"`
	UserName    string `json:"user_name"`
	DecryptData string `json:"decrypt_data"`
}
