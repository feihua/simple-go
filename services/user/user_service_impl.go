package user

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

type UserServiceImpl struct {
}

func (h *UserServiceImpl) CreateUser(dto dto.UserDto) error {
	return dao.CreateUser(dto)
}

func (h *UserServiceImpl) QueryUserList(current int, pageSize int) ([]models.User, int) {
	return dao.QueryUserList(current, pageSize)
}

func (h *UserServiceImpl) UpdateUser(userDto dto.UserDto) error {
	return dao.UpdateUser(userDto)
}

func (h *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	return dao.DeleteUserByIds(ids)
}

func (h *UserServiceImpl) UpdateUserRole(u requests.UserRoleRequest) error {
	return dao.UpdateUserRole(u.UserId, u.Roleid)
}
