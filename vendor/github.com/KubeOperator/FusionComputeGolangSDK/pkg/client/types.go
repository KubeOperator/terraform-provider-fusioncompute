package client

type LoginResponse struct {
	Validity     int      `json:"validity"`
	RightType    int      `json:"rightType"`
	PrivilegeIds []string `json:"privilegeIds"`
	UserId       string   `json:"userId"`
	UserName     string   `json:"userName"`
	RoleList     []string `json:"roleList"`
}
