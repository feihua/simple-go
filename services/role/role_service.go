package role

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

type RoleService interface {
	CreateRole(dto dto.RoleDto) error

	QueryRoleList(current int, pageSize int) ([]models.Role, int)

	UpdateRole(roleDto dto.RoleDto) error

	DeleteRoleByIds(ids []int64) error

	QueryRoleMenuList(id string) []int64
	UpdateRoleMenu(request requests.RoleMenuRequest)
}
