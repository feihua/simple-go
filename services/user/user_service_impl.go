package user

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
	"strconv"
)

// UserServiceImpl 用户操作接口实现
/*
Author: LiuFeiHua
Date: 2024/4/16 13:58
*/
type UserServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewUserServiceImpl(Dao *dao.DaoImpl) UserService {
	return &UserServiceImpl{
		Dao: Dao,
	}
}

// CreateUser 创建用户
func (u *UserServiceImpl) CreateUser(dto dto.UserDto) error {
	return u.Dao.UserDao.CreateUser(dto)
}

// QueryUserList 查询用户列表
func (u *UserServiceImpl) QueryUserList(current int, pageSize int) ([]models.User, int) {
	return u.Dao.UserDao.QueryUserList(current, pageSize)
}

// UpdateUser 更新用户
func (u *UserServiceImpl) UpdateUser(userDto dto.UserDto) error {
	return u.Dao.UserDao.UpdateUser(userDto)
}

// DeleteUserByIds 删除用户
func (u *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	return u.Dao.UserDao.DeleteUserByIds(ids)
}

// QueryUserRoleList 查询用户与角色关糸
func (u *UserServiceImpl) QueryUserRoleList(id string) []int64 {
	userId, _ := strconv.ParseInt(id, 10, 64)
	return u.Dao.UserRoleDao.QueryUserRoleList(userId)
}

// UpdateUserRoleList 更新用户与角色关糸
func (u *UserServiceImpl) UpdateUserRoleList(requests.UserRoleRequest) error {
	return u.Dao.UserRoleDao.UpdateUserRoleList(1, 1)
}
