package user

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

type UserService struct {
}

func (h *UserService) CreateUser(dto dto.UserDto) error {
	return dao.CreateUser(dto)
}

func (h *UserService) QueryUserList(current int, pageSize int) ([]models.User, int) {
	return dao.QueryUserList(current, pageSize)
}

func (h *UserService) UpdateUser(userDto dto.UserDto) error {
	return dao.UpdateUser(userDto)
}

func (h *UserService) DeleteUserById(id int64) error {
	return dao.DeleteUserById(id)
}

func (h *UserService) UpdateUserRole(u requests.UserRoleRequest) error {
	return dao.UpdateUserRole(u.UserId, u.Roleid)
}
