package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

type UserService interface {
	CreateUser(dto dto.UserDto) error

	QueryUserList(current int, pageSize int) ([]models.User, int)

	UpdateUser(userDto dto.UserDto) error

	DeleteUserByIds(ids []int64) error
	UpdateUserRole(u requests.UserRoleRequest) error
}
