package role

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
	"strconv"
)

type RoleService struct {
}

func (r *RoleService) CreateRole(dto dto.RoleDto) error {
	return dao.CreateSysRole(dto)
}

func (r *RoleService) QueryRoleList(current int, pageSize int) ([]models.Role, int) {
	return dao.QueryRoleList(current, pageSize)
}

func (r *RoleService) UpdateRole(roleDto dto.RoleDto) error {
	return dao.UpdateSysRole(roleDto)
}

func (r *RoleService) DeleteRoleById(id int64) error {
	return dao.DeleteSysRoleById(id)
}

func (r *RoleService) QueryRoleMenuList(id string) []int64 {
	roleId, _ := strconv.ParseInt(id, 10, 64)
	return dao.QueryRoleMenuList(roleId)
}

func (r *RoleService) UpdateRoleMenu(request requests.RoleMenuRequest) {
	dao.UpdateRoleMenu(request.RoleId, request.MenuIds)
}
