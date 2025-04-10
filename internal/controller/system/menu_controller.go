package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/menu"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
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

	req := rq.AddMenuReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AddMenuDto{
		MenuName: req.MenuName, // 菜单名称
		MenuType: req.MenuType, // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:  req.Visible,  // 显示状态（0:隐藏, 显示:1）
		Status:   req.Status,   // 菜单状态(1:正常，0:禁用)
		Sort:     req.Sort,     // 排序
		ParentId: req.ParentId, // 父ID
		MenuUrl:  req.MenuUrl,  // 路由路径
		ApiUrl:   req.ApiUrl,   // 接口URL
		MenuIcon: req.MenuIcon, // 菜单图标
		Remark:   req.Remark,   // 备注
		CreateBy: c.MustGet("userName").(string),
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

	req := rq.DeleteMenuReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
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

	req := rq.UpdateMenuReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateMenuDto{
		Id:       req.Id,       // 主键
		MenuName: req.MenuName, // 菜单名称
		MenuType: req.MenuType, // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:  req.Visible,  // 显示状态（0:隐藏, 显示:1）
		Status:   req.Status,   // 菜单状态(1:正常，0:禁用)
		Sort:     req.Sort,     // 排序
		ParentId: req.ParentId, // 父ID
		MenuUrl:  req.MenuUrl,  // 路由路径
		ApiUrl:   req.ApiUrl,   // 接口URL
		MenuIcon: req.MenuIcon, // 菜单图标
		Remark:   req.Remark,   // 备注
		UpdateBy: c.MustGet("userName").(string),
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

	req := rq.UpdateMenuStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateMenuStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
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
	req := rq.QueryMenuDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryMenuDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryMenuDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.OkWithData(c, data)
	}
}

// QueryMenuList 查询菜单信息列表
func (r MenuController) QueryMenuList(c *gin.Context) {
	req := rq.QueryMenuListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryMenuListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}
	list, total, err := r.Service.QueryMenuList(item)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}
