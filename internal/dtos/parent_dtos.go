package dtos

type ParentLoginRequest struct {
	ParentID string `json:"parent_id"`
}

type ParentLoginResponse struct {
	FullName        string `json:"full_name"`
	AccountUsername string `json:"account_username"`
	AccountPassword string `json:"account_password"`
}
