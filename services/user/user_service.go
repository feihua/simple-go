package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/vo/requests"
)

// UserService 用户操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:58
*/
type UserService interface {

	// Login 用户登录
	Login(dto dto.UserLoginDto) (*dto.LoginDtoResp, error)
	// CreateUser 创建用户
	CreateUser(dto dto.UserDto) error
	// QueryUserList 查询用户列表
	QueryUserList(current int, pageSize int) ([]models.User, int)
	// UpdateUser 更新用户
	UpdateUser(userDto dto.UserDto) error
	// DeleteUserByIds 删除用户
	DeleteUserByIds(ids []int64) error
	// QueryUserRoleList 查询用户与角色关糸
	QueryUserRoleList(id string) []int64
	// UpdateUserRoleList 更新用户与角色关糸
	UpdateUserRoleList(u requests.UserRoleRequest) error
}
