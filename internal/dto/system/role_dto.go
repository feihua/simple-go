package system

import "time"

// AddRoleDto 添加角色信息请求参数
type AddRoleDto struct {
	RoleName  string `json:"roleName"`  // 名称
	RoleKey   string `json:"roleKey"`   // 角色权限字符串
	DataScope int32  `json:"dataScope"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    int32  `json:"status"`    // 状态(1:正常，0:禁用)
	Remark    string `json:"remark"`    // 备注
	CreateBy  string `json:"createBy"`  // 创建者

}

// DeleteRoleDto 删除角色信息请求参数
type DeleteRoleDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateRoleDto 修改角色信息请求参数
type UpdateRoleDto struct {
	Id         int64     `json:"id"`         // 主键
	RoleName   string    `json:"roleName"`   // 名称
	RoleKey    string    `json:"roleKey"`    // 角色权限字符串
	DataScope  int32     `json:"dataScope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int32     `json:"status"`     // 状态(1:正常，0:禁用)
	Remark     string    `json:"remark"`     // 备注
	DelFlag    int32     `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string    `json:"createBy"`   // 创建者
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateBy   string    `json:"updateBy"`   // 更新者
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}

// UpdateRoleStatusDto 修改角色信息状态请求参数
type UpdateRoleStatusDto struct {
	Ids      []int64 `json:"ids"`      // id
	Status   int32   `json:"status"`   // 状态（0:关闭,1:正常 ）
	UpdateBy string  `json:"updateBy"` // 更新者
}

// QueryRoleDetailDto 查询角色信息详情请求参数
type QueryRoleDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryRoleListDto 查询角色信息列表请求参数
type QueryRoleListDto struct {
	PageNo    int    `json:"pageNo" default:"1"`    // 第几页
	PageSize  int    `json:"pageSize" default:"10"` // 每页的数量
	RoleName  string `json:"roleName"`              // 名称
	RoleKey   string `json:"roleKey"`               // 角色权限字符串
	DataScope *int32 `json:"dataScope"`             // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    *int32 `json:"status"`                // 状态(1:正常，0:禁用)

}

type QueryRoleMenuListDto struct {
	RoleId int64 `json:"roleId"` // 角色Id
}

type RoleMenuListDataDto struct {
	Key           string `json:"key"`      // 菜单名称
	Title         string `json:"title"`    // 菜单名称
	ParentId      int64  `json:"parentId"` // 父菜单ID，一级菜单为0
	Id            int64  `json:"id"`       // 父菜单ID，一级菜单为0
	Label         string `json:"label"`    // 父菜单ID，一级菜单为0
	IsPenultimate bool   `json:"isPenultimate"`
}
type QueryRoleMenuListDataDtoResp struct {
	MenuList []RoleMenuListDataDto `json:"menuList"`
	MenuIds  []int64               `json:"menuIds"`
}

type UpdateRoleMenuDto struct {
	RoleId  int64   `json:"roleId" binding:"required"`  // 角色Id
	MenuIds []int64 `json:"menuIds" binding:"required"` // 菜单Id
}

// QueryRoleListDtoResp 查询角色信息列表响应参数
type QueryRoleListDtoResp struct {
	Id         int64  `json:"id"`         // 主键
	RoleName   string `json:"roleName"`   // 名称
	RoleKey    string `json:"roleKey"`    // 角色权限字符串
	DataScope  int32  `json:"dataScope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     int32  `json:"status"`     // 状态(1:正常，0:禁用)
	Remark     string `json:"remark"`     // 备注
	DelFlag    int32  `json:"delFlag"`    // 删除标志（0代表删除 1代表存在）
	CreateBy   string `json:"createBy"`   // 创建者
	CreateTime string `json:"createTime"` // 创建时间
	UpdateBy   string `json:"updateBy"`   // 更新者
	UpdateTime string `json:"updateTime"` // 更新时间
}

// QueryRoleUserListDto 查询角色用户信息列表请求参数
type QueryRoleUserListDto struct {
	PageNo   int    `json:"pageNo" `   // 第几页
	PageSize int    `json:"pageSize"`  // 每页的数量
	RoleId   int64  `json:"RoleId" `   // 主键
	Mobile   string `json:"mobile" `   // 手机号码
	UserName string `json:"userName" ` // 用户账号
}

type AuthUserDto struct {
	RoleId int64 `json:"roleId"` // 角色Id
	UserId int64 `json:"userId"` // 用户Id
}
type BatchAuthUserDto struct {
	RoleId  int64   `json:"roleId"`  // 角色Id
	UserIds []int64 `json:"userIds"` // 用户Id
}
