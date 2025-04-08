package user

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// UserServiceImpl 用户信息操作实现
type UserServiceImpl struct {
	Dao *system.UserDao
}

func NewUserServiceImpl(dao *system.UserDao) UserService {
	return &UserServiceImpl{
		Dao: dao,
	}
}

// CreateUser 添加用户信息
func (s *UserServiceImpl) CreateUser(dto a.AddUserDto) error {
	return s.Dao.CreateUser(dto)
}

// DeleteUserByIds 删除用户信息
func (s *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	return s.Dao.DeleteUserByIds(ids)
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(dto a.UpdateUserDto) error {
	return s.Dao.UpdateUser(dto)
}

// UpdateUserStatus 更新用户信息状态
func (s *UserServiceImpl) UpdateUserStatus(dto a.UpdateUserStatusDto) error {
	return s.Dao.UpdateUserStatus(dto)
}

// QueryUserDetail 查询用户信息详情
func (s *UserServiceImpl) QueryUserDetail(dto a.QueryUserDetailDto) (b.User, error) {
	return s.Dao.QueryUserDetail(dto)
}

// QueryUserList 查询用户信息列表
func (s *UserServiceImpl) QueryUserList(dto a.QueryUserListDto) ([]b.User, int64) {
	return s.Dao.QueryUserList(dto)
}
