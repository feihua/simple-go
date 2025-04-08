package role

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// RoleService 角色信息操作接口
type RoleService interface {
	// CreateRole 添加角色信息
	CreateRole(dto b.AddRoleDto) error
	// DeleteRoleByIds 删除角色信息
	DeleteRoleByIds(ids []int64) error
	// UpdateRole 更新角色信息
	UpdateRole(dto b.UpdateRoleDto) error
	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(dto b.UpdateRoleStatusDto) error
	// QueryRoleDetail 查询角色信息详情
	QueryRoleDetail(dto b.QueryRoleDetailDto) (a.Role, error)
	// QueryRoleList 查询角色信息列表
	QueryRoleList(dto b.QueryRoleListDto) ([]a.Role, int64)
}
