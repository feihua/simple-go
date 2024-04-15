package role

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
	"strconv"
)

type RoleServiceImpl struct {
}

func (r *RoleServiceImpl) CreateRole(dto dto.RoleDto) error {
	return dao.CreateSysRole(dto)
}

func (r *RoleServiceImpl) QueryRoleList(current int, pageSize int) ([]models.Role, int) {
	return dao.QueryRoleList(current, pageSize)
}

func (r *RoleServiceImpl) UpdateRole(roleDto dto.RoleDto) error {
	return dao.UpdateSysRole(roleDto)
}

func (r *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return dao.DeleteSysRoleByIds(ids)
}

func (r *RoleServiceImpl) QueryRoleMenuList(id string) []int64 {
	roleId, _ := strconv.ParseInt(id, 10, 64)
	return dao.QueryRoleMenuList(roleId)
}

func (r *RoleServiceImpl) UpdateRoleMenu(request requests.RoleMenuRequest) {
	dao.UpdateRoleMenu(request.RoleId, request.MenuIds)
}
