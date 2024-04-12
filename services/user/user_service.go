package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
	"github.com/feihua/simple-go/requests"
)

type UserService struct {
}

func (h *UserService) CreateUser(dto dto.UserDto) error {
	return repositories.CreateUser(dto)
}

func (h *UserService) GetUserList(current int, pageSize int) ([]models.SysUser, int) {
	return repositories.GetUserList(current, pageSize)
}

func (h *UserService) UpdateUser(userDto dto.UserDto) error {
	return repositories.UpdateUser(userDto)
}

func (h *UserService) DeleteUserById(id int64) error {
	return repositories.DeleteUserById(id)
}

func (h *UserService) UpdateUserRole(u requests.UserRoleRequest) error {
	return repositories.UpdateUserRole(u.UserId, u.Roleid)
}
