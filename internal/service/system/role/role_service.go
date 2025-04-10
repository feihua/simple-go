package role

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// RoleService 角色信息操作接口
type RoleService interface {
	// CreateRole 添加角色信息
	CreateRole(dto d.AddRoleDto) error
	// DeleteRoleByIds 删除角色信息
	DeleteRoleByIds(ids []int64) error
	// UpdateRole 更新角色信息
	UpdateRole(dto d.UpdateRoleDto) error
	// UpdateRoleStatus 更新角色信息状态
	UpdateRoleStatus(dto d.UpdateRoleStatusDto) error
	// QueryRoleDetail 查询角色信息详情
	QueryRoleDetail(dto d.QueryRoleDetailDto) (*d.QueryRoleListDtoResp, error)
	// QueryRoleById 根据id查询角色信息详情
	QueryRoleById(id int64) (*d.QueryRoleListDtoResp, error)
	// QueryRoleList 查询角色信息列表
	QueryRoleList(dto d.QueryRoleListDto) ([]*d.QueryRoleListDtoResp, int64, error)
}
