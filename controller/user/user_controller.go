package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
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

	//校验通过，返回请求参数
	c.JSON(http.StatusOK, gin.H{"status": "ok", "type": "account", "currentAuthority": "admin"})
}

// QueryUserInfo 查询用户信息
func (u UserController) QueryUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"name": "liufeihua", "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"})
}

// CreateUser 创建用户
func (u UserController) CreateUser(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	userDto := dto.UserDto{
		Mobile:   userRequest.Mobile,
		UserName: userRequest.UserName,
		Password: userRequest.Password,
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
	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	list, total := u.Service.UserService.QueryUserList(pageNum, size)

	result.OkWithData(c, gin.H{"list": list, "success": true, "current": current, "total": total, "pageSize": pageSize})
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

// QueryUserRoleList 查询用户与角色关糸
func (u UserController) QueryUserRoleList(c *gin.Context) {
	userId := c.DefaultQuery("userId", "1")

	menuIds := u.Service.UserService.QueryUserRoleList(userId)
	c.JSON(http.StatusOK, gin.H{"data": menuIds})
}

// UpdateUserRoleList 更新用户与角色关糸
func (u UserController) UpdateUserRoleList(c *gin.Context) {

	req := requests.UserRoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = u.Service.UserService.UpdateUserRoleList(req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}
