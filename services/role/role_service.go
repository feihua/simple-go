package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
	"github.com/feihua/simple-go/requests"
	"strconv"
)

type RoleService struct {
}

func (r *RoleService) CreateRole(dto dto.RoleDto) error {
	return repositories.CreateSysRole(dto)
}

func (r *RoleService) QueryRoleList(current int, pageSize int) ([]models.Role, int) {
	return repositories.QueryRoleList(current, pageSize)
}

func (r *RoleService) UpdateRole(roleDto dto.RoleDto) error {
	return repositories.UpdateSysRole(roleDto)
}

func (r *RoleService) DeleteRoleById(id int64) error {
	return repositories.DeleteSysRoleById(id)
}

func (r *RoleService) QueryRoleMenuList(id string) []int64 {
	roleId, _ := strconv.ParseInt(id, 10, 64)
	return repositories.QueryRoleMenuList(roleId)
}

func (r *RoleService) UpdateRoleMenu(request requests.RoleMenuRequest) {
	repositories.UpdateRoleMenu(request.RoleId, request.MenuIds)
}
