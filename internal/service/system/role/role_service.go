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
	// QueryAllocatedList 查询已分配用户角色列表
	QueryAllocatedList(item d.QueryRoleUserListDto) ([]*d.QueryUserListDtoResp, int64, error)
	// QueryUnallocatedList 查询未分配用户角色列表
	QueryUnallocatedList(item d.QueryRoleUserListDto) ([]*d.QueryUserListDtoResp, int64, error)
	// CancelAuthUser 取消授权用户
	CancelAuthUser(item d.AuthUserDto) error
	// BatchCancelAuthUser 批量取消授权用户
	BatchCancelAuthUser(item d.BatchAuthUserDto) error
	// BatchAuthUser 批量选择用户授权
	BatchAuthUser(dto d.BatchAuthUserDto) error
	// QueryRoleMenuList 分页查询角色菜单关联列表
	QueryRoleMenuList(item d.QueryRoleMenuListDto) (*d.QueryRoleMenuListDataDtoResp, error)
	// UpdateRoleMenu 添加角色菜单关联
	UpdateRoleMenu(item d.UpdateRoleMenuDto) error
}
