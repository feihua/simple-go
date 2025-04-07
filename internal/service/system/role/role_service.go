package role

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// RoleService 角色操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:47
*/
type RoleService interface {
	// CreateRole 创建角色
	CreateRole(dto system.RoleDto) error
	// QueryRoleList 查询角色列表
	QueryRoleList(role system.QueryRoleListDto) ([]system2.Role, int64)
	// UpdateRole 更新角色
	UpdateRole(roleDto system.RoleDto) error
	// DeleteRoleByIds 删除角色
	DeleteRoleByIds(ids []int64) error
	// QueryRoleMenuList 根据角色Id查询角色菜单
	QueryRoleMenuList(id int64) (*system.QueryRoleMenuListDtoResp, error)
	// UpdateRoleMenuList 更新角色菜单
	UpdateRoleMenuList(dtoRequest system.RoleMenuDtoRequest) error
}
