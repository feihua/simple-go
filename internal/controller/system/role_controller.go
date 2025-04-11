package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/role"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// RoleController 角色信息相关
type RoleController struct {
	Service role.RoleService
}

func NewRoleController(Service role.RoleService) *RoleController {
	return &RoleController{Service: Service}
}

// CreateRole 添加角色信息
func (r RoleController) CreateRole(c *gin.Context) {

	req := rq.AddRoleReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AddRoleDto{
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
		Remark:    req.Remark,    // 备注
		CreateBy:  c.MustGet("userName").(string),
	}

	err = r.Service.CreateRole(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteRoleByIds 删除角色信息
func (r RoleController) DeleteRoleByIds(c *gin.Context) {

	req := rq.DeleteRoleReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteRoleByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateRole 更新角色信息
func (r RoleController) UpdateRole(c *gin.Context) {

	req := rq.UpdateRoleReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateRoleDto{
		Id:        req.Id,        // 主键
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
		Remark:    req.Remark,    // 备注
		UpdateBy:  c.MustGet("userName").(string),
	}
	err = r.Service.UpdateRole(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateRoleStatus 更新角色信息状态
func (r RoleController) UpdateRoleStatus(c *gin.Context) {

	req := rq.UpdateRoleStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateRoleStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateRoleStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryRoleDetail 查询角色信息详情
func (r RoleController) QueryRoleDetail(c *gin.Context) {
	req := rq.QueryRoleDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryRoleDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryRoleDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, data)
	}
}

// QueryRoleList 查询角色信息列表
func (r RoleController) QueryRoleList(c *gin.Context) {
	req := rq.QueryRoleListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryRoleListDto{
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
		RoleName:  req.RoleName,  // 名称
		RoleKey:   req.RoleKey,   // 角色权限字符串
		DataScope: req.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    req.Status,    // 状态(1:正常，0:禁用)
	}
	list, total, err := r.Service.QueryRoleList(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}

// QueryRoleMenuList 分页查询角色菜单关联列表
func (r RoleController) QueryRoleMenuList(c *gin.Context) {
	req := rq.QueryRoleMenuListReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryRoleMenuListDto{
		RoleId: req.RoleId,
	}
	data, err := r.Service.QueryRoleMenuList(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, data)
	}
}

// AddRoleMenu 添加角色菜单关联
func (r RoleController) AddRoleMenu(c *gin.Context) {
	req := rq.UpdateRoleMenuReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateRoleMenuDto{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}
	err = r.Service.UpdateRoleMenu(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryAllocatedList 查询已分配用户角色列表
func (r RoleController) QueryAllocatedList(c *gin.Context) {
	req := rq.QueryRoleUserListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryRoleUserListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		RoleId:   req.RoleId,   // 角色ID
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
	}
	list, total, err := r.Service.QueryAllocatedList(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}

// QueryUnallocatedList 查询未分配用户角色列表
func (r RoleController) QueryUnallocatedList(c *gin.Context) {
	req := rq.QueryRoleUserListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryRoleUserListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		RoleId:   req.RoleId,   // 角色ID
		Mobile:   req.Mobile,   // 手机号码
		UserName: req.UserName, // 用户账号
	}
	list, total, err := r.Service.QueryUnallocatedList(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}

// CancelAuthUser 取消授权用户
func (r RoleController) CancelAuthUser(c *gin.Context) {
	req := rq.AuthUserReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AuthUserDto{
		RoleId: req.RoleId, // 角色ID
		UserId: req.UserId, // 需要授权的用户数据ID集合
	}
	err = r.Service.CancelAuthUser(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// BatchCancelAuthUser 批量取消授权用户
func (r RoleController) BatchCancelAuthUser(c *gin.Context) {
	req := rq.BatchAuthUserReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.BatchAuthUserDto{
		RoleId:  req.RoleId,  // 角色ID
		UserIds: req.UserIds, // 需要授权的用户数据ID集合
	}
	err = r.Service.BatchCancelAuthUser(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}

// BatchAuthUser 批量选择用户授权
func (r RoleController) BatchAuthUser(c *gin.Context) {
	req := rq.BatchAuthUserReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.BatchAuthUserDto{
		RoleId:  req.RoleId,  // 角色ID
		UserIds: req.UserIds, // 需要授权的用户数据ID集合
	}
	err = r.Service.BatchAuthUser(item)
	if err != nil {
		result.FailWithMsg(c, result.RoleError, err.Error())
	} else {
		result.Ok(c)
	}
}
