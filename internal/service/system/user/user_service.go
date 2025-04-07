package user

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// UserService 用户操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:58
*/
type UserService interface {

	// Login 用户登录
	Login(dto system.UserLoginDto) (*system.LoginDtoResp, error)
	// QueryUserMenu 查询用户菜单权限信息
	QueryUserMenu(userId int64, userName string) (*system.QueryUserMenuDtoResp, error)
	// CreateUser 创建用户
	CreateUser(dto system.UserDto) error
	// QueryUserList 查询用户列表
	QueryUserList(userListDto system.QueryUserListDto) ([]system2.User, int64)
	// UpdateUser 更新用户
	UpdateUser(userDto system.UserDto) error
	// DeleteUserByIds 删除用户
	DeleteUserByIds(ids []int64) error
	// QueryUserRoleList 根据用户id查询用户与角色关糸
	QueryUserRoleList(id int64) (*system.QueryUserRoleListDtoResp, error)
	// UpdateUserRoleList 更新用户与角色关糸
	UpdateUserRoleList(u system.UpdateUserRoleDtoRequest) error
}
