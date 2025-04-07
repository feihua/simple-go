package requests

type RoleRequest struct {
	Id       int64  `json:"id"`       //主键
	RoleName string `json:"roleName"` //名称
	StatusId int32  `json:"statusId"` //状态(1:正常，0:禁用)
	Sort     int32  `json:"sort"`     //排序
	Remark   string `json:"remark"`   //备注
}

type RoleMenuRequest struct {
	RoleId  int64   `json:"roleId" binding:"required"`
	MenuIds []int64 `json:"menuIds" binding:"required"`
}
type DeleteRoleRequest struct {
	Ids []int64 `json:"ids" binding:"required"` //编号
}

// QueryRoleListRequest 查询角色列表
type QueryRoleListRequest struct {
	RoleName string `json:"roleName"`                    //名称
	StatusId int32  `json:"statusId" binding:"required"` //状态(1:正常，0:禁用)
	PageNo   int    `json:"pageNo" default:"1"`          //第几页
	PageSize int    `json:"pageSize" default:"10"`       //每页的数量
}
