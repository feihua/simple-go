package user

import (
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/services"
	"github.com/feihua/simple-go/internal/vo/requests"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserController 用户相关
/*
Author: LiuFeiHua
Date: 2024/4/15 16:53
*/
type UserController struct {
	Service *services.ServiceImpl
}

func NewUserController(Service *services.ServiceImpl) *UserController {
	return &UserController{
		Service: Service,
	}
}

// Login 登录
func (u UserController) Login(c *gin.Context) {
	loginRequest := requests.LoginRequest{}
	err := c.ShouldBind(&loginRequest)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	loginDto := dto.UserLoginDto{
		Account:  loginRequest.Account,
		Password: loginRequest.Password,
	}

	loginDtoResp, err := u.Service.UserService.Login(loginDto)
	if err != nil {
		result.FailWithMsg(c, result.UserLoginError, err.Error())
	} else {
		result.OkWithData(c, loginDtoResp)
	}
}

// QueryUserMenuList 查询用户菜单权限信息
func (u UserController) QueryUserMenuList(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	userName := c.MustGet("userName").(string)
	resp, err := u.Service.UserService.QueryUserMenu(int64(userId), userName)

	if err != nil {
		result.FailWithMsg(c, result.UserLoginError, err.Error())
	} else {
		result.OkWithData(c, resp)
	}
}

// CreateUser 创建用户
func (u UserController) CreateUser(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		utils.Logger.Error(err.Error())
		result.Fail(c, result.ParamsError)
		return
	}

	userDto := dto.UserDto{
		Mobile:   userRequest.Mobile,
		UserName: userRequest.UserName,
		Password: "123456",
		StatusId: userRequest.StatusId,
		Sort:     userRequest.Sort,
		Remark:   userRequest.Remark,
	}
	err = u.Service.UserService.CreateUser(userDto)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryUserList 查询用户列表
func (u UserController) QueryUserList(c *gin.Context) {
	req := requests.QueryUserListRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	queryUserListDto := dto.QueryUserListDto{
		Mobile:   req.Mobile,
		UserName: req.UserName,
		StatusId: req.StatusId,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}

	list, total := u.Service.UserService.QueryUserList(queryUserListDto)

	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}

// UpdateUser 更新用户
func (u UserController) UpdateUser(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	userDto := dto.UserDto{
		Id:       userRequest.Id,
		Mobile:   userRequest.Mobile,
		UserName: userRequest.UserName,
		Password: userRequest.Password,
		StatusId: userRequest.StatusId,
		Sort:     userRequest.Sort,
		Remark:   userRequest.Remark,
	}
	err = u.Service.UserService.UpdateUser(userDto)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteUserByIds 删除用户
func (u UserController) DeleteUserByIds(c *gin.Context) {

	userRequest := requests.DeleteUserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = u.Service.UserService.DeleteUserByIds(userRequest.Ids)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryUserRoleList 根据用户id查询用户与角色关糸
func (u UserController) QueryUserRoleList(c *gin.Context) {
	userId := c.DefaultQuery("userId", "10")
	id, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	resp, err := u.Service.UserService.QueryUserRoleList(id)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.OkWithData(c, resp)
	}
}

// UpdateUserRoleList 更新用户与角色关糸
func (u UserController) UpdateUserRoleList(c *gin.Context) {
	req := requests.UpdateUserRoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	roleDtoRequest := dto.UpdateUserRoleDtoRequest{
		UserId: req.UserId,
		RoleId: req.RoleId,
	}

	err = u.Service.UserService.UpdateUserRoleList(roleDtoRequest)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}
