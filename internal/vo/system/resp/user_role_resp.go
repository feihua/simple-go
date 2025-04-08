package system

import "time"

// UserRole 角色用户关联
type UserRole struct {
	Id     int64 `json:"id"`     //主键
	UserId int64 `json:"userId"` //用户ID
	RoleId int64 `json:"roleId"` //角色ID
}
