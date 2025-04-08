package system

import "time"

// AddRoleMenuReqVo 添加菜单角色关联请求参数
type AddRoleMenuReqVo struct {
	Id     int64 `json:"id"`     //主键
	RoleId int64 `json:"roleId"` //角色ID
	MenuId int64 `json:"menuId"` //菜单ID
}

// DeleteRoleMenuReqVo 删除菜单角色关联请求参数
type DeleteRoleMenuReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateRoleMenuReqVo 修改菜单角色关联请求参数
type UpdateRoleMenuReqVo struct {
	Id     int64 `json:"id"`     //主键
	RoleId int64 `json:"roleId"` //角色ID
	MenuId int64 `json:"menuId"` //菜单ID
}

// UpdateRoleMenuStatusReqVo 修改菜单角色关联状态请求参数
type UpdateRoleMenuStatusReqVo struct {
	Ids    []int64 `json:"ids"`    //id
	Status int32   `json:"status"` //状态（0:关闭,1:正常 ）
}

// QueryRoleMenuDetailReqVo 查询菜单角色关联详情请求参数
type QueryRoleMenuDetailReqVo struct {
	Id int64 `json:"id"` //id
}

// QueryRoleMenuListReqVo 查询菜单角色关联列表请求参数
type QueryRoleMenuListReqVo struct {
	PageNo   int   `json:"pageNo" default:"1"`    //第几页
	PageSize int   `json:"pageSize" default:"10"` //每页的数量
	Id       int64 `json:"id"`                    //主键
	RoleId   int64 `json:"roleId"`                //角色ID
	MenuId   int64 `json:"menuId"`                //菜单ID
}
