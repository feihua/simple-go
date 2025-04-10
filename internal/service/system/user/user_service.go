package user

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// UserService 用户信息操作接口
type UserService interface {
	// CreateUser 添加用户信息
	CreateUser(dto d.AddUserDto) error
	// DeleteUserByIds 删除用户信息
	DeleteUserByIds(ids []int64) error
	// UpdateUser 更新用户信息
	UpdateUser(dto d.UpdateUserDto) error
	// UpdateUserStatus 更新用户信息状态
	UpdateUserStatus(dto d.UpdateUserStatusDto) error
	// QueryUserDetail 查询用户信息详情
	QueryUserDetail(dto d.QueryUserDetailDto) (*d.QueryUserListDtoResp, error)
	// QueryUserList 查询用户信息列表
	QueryUserList(dto d.QueryUserListDto) ([]*d.QueryUserListDtoResp, int64, error)

	// Login 用户登录
	Login(dto d.LoginDto) (*string, error)
	// QueryUserMenu 查询用户菜单权限信息
	QueryUserMenu(userId int64, userName string) (*d.QueryUserMenuDtoResp, error)
}
