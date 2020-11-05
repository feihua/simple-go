package requests

type RoleRequest struct {
	ID       int64  `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RoleMenuRequest struct {
	RoleId  int64   `json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}
