package role

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
	"strconv"
)

// RoleServiceImpl 角色操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 13:47
*/
type RoleServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewRoleServiceImpl(dao *dao.DaoImpl) RoleService {
	return &RoleServiceImpl{Dao: dao}
}

// CreateRole 创建角色
func (r *RoleServiceImpl) CreateRole(dto dto.RoleDto) error {
	return r.Dao.RoleDao.CreateRole(dto)
}

// QueryRoleList 查询角色列表
func (r *RoleServiceImpl) QueryRoleList(current int, pageSize int) ([]models.Role, int64) {
	return r.Dao.RoleDao.QueryRoleList(current, pageSize)
}

// UpdateRole 更新角色
func (r *RoleServiceImpl) UpdateRole(roleDto dto.RoleDto) error {
	return r.Dao.RoleDao.UpdateRole(roleDto)
}

// DeleteRoleByIds 删除角色
func (r *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return r.Dao.RoleDao.DeleteRoleByIds(ids)
}

// QueryRoleMenuList 查询角色菜单
func (r *RoleServiceImpl) QueryRoleMenuList(id string) []int64 {
	roleId, _ := strconv.ParseInt(id, 10, 64)
	return r.Dao.RoleMenuDao.QueryRoleMenuList(roleId)
}

// UpdateRoleMenuList 更新角色菜单
func (r *RoleServiceImpl) UpdateRoleMenuList(request requests.RoleMenuRequest) {
	r.Dao.RoleMenuDao.UpdateRoleMenuList(request.RoleId, request.MenuIds)
}
