package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

type UserContract interface {
	CreateUser(dto dto.UserDto) error

	QueryUserList(current int, pageSize int) ([]models.User, int)

	UpdateUser(userDto dto.UserDto) error

	DeleteUserById(id int64) error
	UpdateUserRole(u requests.UserRoleRequest) error
}
