package user

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// UserService 用户信息操作接口
type UserService interface {
	// CreateUser 添加用户信息
	CreateUser(dto b.AddUserDto) error
	// DeleteUserByIds 删除用户信息
	DeleteUserByIds(ids []int64) error
	// UpdateUser 更新用户信息
	UpdateUser(dto b.UpdateUserDto) error
	// UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(dto b.UpdateUserStatusDto) error
	// QueryUserDetail 查询用户信息详情
	QueryUserDetail(dto b.QueryUserDetailDto) (a.User, error)
	// QueryUserList 查询用户信息列表
	QueryUserList(dto b.QueryUserListDto) ([]a.User, int64)

	// Login 用户登录
	Login(dto b.LoginDto) (*b.LoginDtoResp, error)
}
