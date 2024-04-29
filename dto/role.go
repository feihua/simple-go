package dto

type RoleDto struct {
	Id       int64  `json:"id"`       //主键
	RoleName string `json:"roleName"` //名称
	StatusId int8   `json:"statusId"` //状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     //排序
	Remark   string `json:"remark"`   //备注
}
type RoleMenuDtoRequest struct {
	RoleId  int64   `json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}

// QueryRoleListDto 查询角色列表
type QueryRoleListDto struct {
	RoleName string `json:"roleName"` //名称
	StatusId int8   `json:"statusId"` //状态(1:正常，0:禁用)
	PageNo   int    `json:"pageNo"`   //第几页
	PageSize int    `json:"pageSize"` //每页的数量
}
