package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/menu"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// MenuController 菜单信息相关
type MenuController struct {
	Service menu.MenuService
}

func NewMenuController(Service menu.MenuService) *MenuController {
	return &MenuController{Service: Service}
}

// CreateMenu 添加菜单信息
func (r MenuController) CreateMenu(c *gin.Context) {

	req := b.AddMenuReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.AddMenuDto{
		Id:         req.Id,         // 主键
		MenuName:   req.MenuName,   // 菜单名称
		MenuType:   req.MenuType,   // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    req.Visible,    // 显示状态（0:隐藏, 显示:1）
		Status:     req.Status,     // 菜单状态(1:正常，0:禁用)
		Sort:       req.Sort,       // 排序
		ParentId:   req.ParentId,   // 父ID
		MenuUrl:    req.MenuUrl,    // 路由路径
		ApiUrl:     req.ApiUrl,     // 接口URL
		MenuIcon:   req.MenuIcon,   // 菜单图标
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}

	err = r.Service.CreateMenu(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteMenuByIds 删除菜单信息
func (r MenuController) DeleteMenuByIds(c *gin.Context) {

	req := b.DeleteMenuReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = r.Service.DeleteMenuByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateMenu 更新菜单信息
func (r MenuController) UpdateMenu(c *gin.Context) {

	req := b.UpdateMenuReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateMenuDto{
		Id:         req.Id,         // 主键
		MenuName:   req.MenuName,   // 菜单名称
		MenuType:   req.MenuType,   // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    req.Visible,    // 显示状态（0:隐藏, 显示:1）
		Status:     req.Status,     // 菜单状态(1:正常，0:禁用)
		Sort:       req.Sort,       // 排序
		ParentId:   req.ParentId,   // 父ID
		MenuUrl:    req.MenuUrl,    // 路由路径
		ApiUrl:     req.ApiUrl,     // 接口URL
		MenuIcon:   req.MenuIcon,   // 菜单图标
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}
	err = r.Service.UpdateMenu(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateMenuStatus 更新菜单信息状态
func (r MenuController) UpdateMenuStatus(c *gin.Context) {

	req := b.UpdateMenuStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateMenuStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateMenuStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryMenuDetail 查询菜单信息详情
func (r MenuController) QueryMenuDetail(c *gin.Context) {
	req := b.QueryMenuDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryMenuDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryMenuDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryMenuList 查询菜单信息列表
func (r MenuController) QueryMenuList(c *gin.Context) {
	req := b.QueryMenuListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryMenuListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		MenuName: req.MenuName, // 菜单名称
		MenuType: req.MenuType, // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:  req.Visible,  // 显示状态（0:隐藏, 显示:1）
		Status:   req.Status,   // 菜单状态(1:正常，0:禁用)
		ParentId: req.ParentId, // 父ID
		MenuUrl:  req.MenuUrl,  // 路由路径
		ApiUrl:   req.ApiUrl,   // 接口URL
		MenuIcon: req.MenuIcon, // 菜单图标
	}
	list, total := r.Service.QueryMenuList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
