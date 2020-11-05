package user

import (
	"simple-go/dto"
	"simple-go/models"
	"simple-go/requests"
)

type UserContract interface {

	CreateUser(dto dto.UserDto) error

	GetUserList(current int, pageSize int) ([]models.SysUser, int)

	UpdateUser(userDto dto.UserDto) error

	DeleteUserById(id int64) error
	UpdateUserRole(u requests.UserRoleRequest) error
}
