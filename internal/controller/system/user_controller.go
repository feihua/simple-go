package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/user"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// UserController 用户信息相关
type UserController struct {
	Service user.UserService
}

func NewUserController(Service user.UserService) *UserController {
	return &UserController{Service: Service}
}

// CreateUser 添加用户信息
func (r UserController) CreateUser(c *gin.Context) {

	req := rq.AddUserReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AddUserDto{
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
		NickName: req.NickName, // 用户昵称
		UserType: req.UserType, // 用户类型（00系统用户）
		Avatar:   req.Avatar,   // 头像路径
		Email:    req.Email,    // 用户邮箱
		Password: req.Password, // 密码
		Status:   req.Status,   // 状态(1:正常，0:禁用)
		DeptId:   req.DeptId,   // 部门ID
		Remark:   req.Remark,   // 备注
		CreateBy: c.MustGet("userName").(string),
	}

	err = r.Service.CreateUser(item)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteUserByIds 删除用户信息
func (r UserController) DeleteUserByIds(c *gin.Context) {

	req := rq.DeleteUserReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteUserByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateUser 更新用户信息
func (r UserController) UpdateUser(c *gin.Context) {

	req := rq.UpdateUserReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateUserDto{
		Id:       req.Id,       // 主键
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
		NickName: req.NickName, // 用户昵称
		UserType: req.UserType, // 用户类型（00系统用户）
		Avatar:   req.Avatar,   // 头像路径
		Email:    req.Email,    // 用户邮箱
		Status:   req.Status,   // 状态(1:正常，0:禁用)
		DeptId:   req.DeptId,   // 部门ID
		Remark:   req.Remark,   // 备注
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateUser(item)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateUserStatus 更新用户信息状态
func (r UserController) UpdateUserStatus(c *gin.Context) {

	req := rq.UpdateUserStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateUserStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateUserStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryUserDetail 查询用户信息详情
func (r UserController) QueryUserDetail(c *gin.Context) {
	req := rq.QueryUserDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryUserDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryUserDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryUserList 查询用户信息列表
func (r UserController) QueryUserList(c *gin.Context) {
	req := rq.QueryUserListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryUserListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
		NickName: req.NickName, // 用户昵称
		Status:   req.Status,   // 状态(1:正常，0:禁用)
		DeptId:   req.DeptId,   // 部门ID

	}
	list, total, err := r.Service.QueryUserList(item)
	if err != nil {
		result.FailWithMsg(c, result.UserError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}

// Login 登录
func (u UserController) Login(c *gin.Context) {
	loginReq := rq.LoginReqVo{}
	err := c.ShouldBind(&loginReq)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	loginDto := d.LoginDto{
		Account:  loginReq.Account,
		Password: loginReq.Password,
	}

	token, err := u.Service.Login(loginDto)
	if err != nil {
		result.FailWithMsg(c, result.UserLoginError, err.Error())
	} else {
		result.OkWithData(c, token)
	}
}

// QueryUserMenuList 查询用户菜单权限信息
func (u UserController) QueryUserMenuList(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	userName := c.MustGet("userName").(string)
	resp, err := u.Service.QueryUserMenu(int64(userId), userName)

	if err != nil {
		result.FailWithMsg(c, result.UserLoginError, err.Error())
	} else {
		result.OkWithData(c, resp)
	}
}

//
// // QueryUserRoleList 根据用户id查询用户与角色关糸
// func (u UserController) QueryUserRoleList(c *gin.Context) {
// 	userId := c.DefaultQuery("userId", "10")
// 	id, err := strconv.ParseInt(userId, 10, 64)
//
// 	if err != nil {
// 		result.Fail(c, result.ParamsError)
// 		return
// 	}
//
// 	resp, err := u.Service.QueryUserRoleList(id)
// 	if err != nil {
// 		result.FailWithMsg(c, result.UserError, err.Error())
// 	} else {
// 		result.OkWithData(c, resp)
// 	}
// }
//
// // UpdateUserRoleList 更新用户与角色关糸
// func (u UserController) UpdateUserRoleList(c *gin.Context) {
// 	req := requests.UpdateUserRoleRequest{}
// 	err := c.ShouldBind(&req)
// 	if err != nil {
// 		result.Fail(c, result.ParamsError)
// 		return
// 	}
//
// 	roleDtoRequest := system.UpdateUserRoleDtoRequest{
// 		UserId: req.UserId,
// 		RoleId: req.RoleId,
// 	}
//
// 	err = u.Service.UpdateUserRoleList(roleDtoRequest)
// 	if err != nil {
// 		result.FailWithMsg(c, result.UserError, err.Error())
// 	} else {
// 		result.Ok(c)
// 	}
// }
