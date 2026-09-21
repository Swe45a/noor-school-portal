package dtos

type TeacherLoginRequest struct {
	IDNumber       string `json:"id_number"`
	EmployeeNumber string `json:"employee_number"`
}

type TeacherLoginResponse struct {
	FullName        string `json:"full_name"`
	AccountUsername string `json:"account_username"`
	AccountPassword string `json:"account_password"`
}
