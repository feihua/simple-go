package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

// RoleService 角色操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:47
*/
type RoleService interface {
	// CreateRole 创建角色
	CreateRole(dto dto.RoleDto) error
	// QueryRoleList 查询角色列表
	QueryRoleList(current int, pageSize int) ([]models.Role, int64)
	// UpdateRole 更新角色
	UpdateRole(roleDto dto.RoleDto) error
	// DeleteRoleByIds 删除角色
	DeleteRoleByIds(ids []int64) error
	// QueryRoleMenuList 查询角色菜单
	QueryRoleMenuList(id string) []int64
	// UpdateRoleMenuList 更新角色菜单
	UpdateRoleMenuList(request requests.RoleMenuRequest)
}
