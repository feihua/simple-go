package menu

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MenuController 菜单相关
/*
Author: LiuFeiHua
Date: 2024/4/15 18:03
*/
type MenuController struct {
	Service *services.ServiceImpl
}

func NewMenuController(Service *services.ServiceImpl) *MenuController {
	return &MenuController{Service: Service}
}

// CreateMenu 创建菜单
func (m MenuController) CreateMenu(c *gin.Context) {

	req := requests.MenuRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	menuDto := dto.MenuDto{
		MenuName: req.MenuName,
		MenuType: req.MenuType,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		ParentId: req.ParentId,
		MenuUrl:  req.MenuUrl,
		ApiUrl:   req.ApiUrl,
		MenuIcon: req.MenuIcon,
		Remark:   req.Remark,
	}

	err = m.Service.MenuService.CreateMenu(menuDto)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryMenuList 查询菜单
func (m MenuController) QueryMenuList(c *gin.Context) {
	req := requests.MenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	menuList, _ := m.Service.MenuService.QueryMenuList("u")
	c.JSON(http.StatusOK, gin.H{"data": menuList})
}

// UpdateMenu 更新菜单
func (m MenuController) UpdateMenu(c *gin.Context) {

	req := requests.MenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	menuDto := dto.MenuDto{
		Id:       req.Id,
		MenuName: req.MenuName,
		MenuType: req.MenuType,
		StatusId: req.StatusId,
		Sort:     req.Sort,
		ParentId: req.ParentId,
		MenuUrl:  req.MenuUrl,
		ApiUrl:   req.ApiUrl,
		MenuIcon: req.MenuIcon,
		Remark:   req.Remark,
	}
	err = m.Service.MenuService.UpdateMenu(menuDto)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteMenuById 删除菜单
func (m MenuController) DeleteMenuById(c *gin.Context) {

	req := requests.DeleteMenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = m.Service.MenuService.DeleteMenuById(req.Id)
	if err != nil {
		result.FailWithMsg(c, result.MenuError, err.Error())
	} else {
		result.Ok(c)
	}
}
