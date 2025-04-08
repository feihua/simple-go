package system

import "time"

// RoleDept 角色和部门关联
type RoleDept struct {
	RoleId int64 `json:"roleId"` //角色id
	DeptId int64 `json:"deptId"` //部门id
}
