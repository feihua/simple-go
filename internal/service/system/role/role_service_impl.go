package role

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// RoleServiceImpl 角色信息操作实现
type RoleServiceImpl struct {
	Dao *system.RoleDao
}

func NewRoleServiceImpl(dao *system.RoleDao) RoleService {
	return &RoleServiceImpl{
		Dao: dao,
	}
}

// CreateRole 添加角色信息
func (s *RoleServiceImpl) CreateRole(dto a.AddRoleDto) error {
	return s.Dao.CreateRole(dto)
}

// DeleteRoleByIds 删除角色信息
func (s *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return s.Dao.DeleteRoleByIds(ids)
}

// UpdateRole 更新角色信息
func (s *RoleServiceImpl) UpdateRole(dto a.UpdateRoleDto) error {
	return s.Dao.UpdateRole(dto)
}

// UpdateRoleStatus 更新角色信息状态
func (s *RoleServiceImpl) UpdateRoleStatus(dto a.UpdateRoleStatusDto) error {
	return s.Dao.UpdateRoleStatus(dto)
}

// QueryRoleDetail 查询角色信息详情
func (s *RoleServiceImpl) QueryRoleDetail(dto a.QueryRoleDetailDto) (b.Role, error) {
	return s.Dao.QueryRoleDetail(dto)
}

// QueryRoleList 查询角色信息列表
func (s *RoleServiceImpl) QueryRoleList(dto a.QueryRoleListDto) ([]b.Role, int64) {
	return s.Dao.QueryRoleList(dto)
}
