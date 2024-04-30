package requests

type RoleRequest struct {
	Id       int64  `json:"id"`       //主键
	RoleName string `json:"roleName"` //名称
	StatusId int32  `json:"statusId"` //状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     //排序
	Remark   string `json:"remark"`   //备注
}

type RoleMenuRequest struct {
	RoleId  int64   `json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}
type DeleteRoleRequest struct {
	Ids []int64 `json:"ids"` //编号
}
