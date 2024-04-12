package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/requests"
)

type RoleContract interface {
	CreateRole(dto dto.RoleDto) error

	GetRoleList(current int, pageSize int) ([]models.SysRole, int)

	UpdateRole(roleDto dto.RoleDto) error

	DeleteRoleById(id int64) error

	GeRoleMenuList(id string) []int64
	UpdateRoleMenu(request requests.RoleMenuRequest)
}
