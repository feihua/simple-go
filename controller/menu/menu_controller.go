package menu

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/menu"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MenuController struct {
}

func NewMenuController() *MenuController {
	return &MenuController{}
}
func (MenuController) CreateMenu(c *gin.Context) {

	req := requests.MenuRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuService = &menu.MenuServiceImpl{}

	u := dto.MenuDto{
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
	result := service.CreateMenu(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (MenuController) QueryMenuList(c *gin.Context) {
	req := requests.MenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuService = &menu.MenuServiceImpl{}

	menuList, _ := service.QueryMenuList("u")
	c.JSON(http.StatusOK, gin.H{"data": menuList})
}

func (MenuController) UpdateMenu(c *gin.Context) {

	req := requests.MenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuService = &menu.MenuServiceImpl{}
	u := dto.MenuDto{
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
	result := service.UpdateMenu(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (MenuController) DeleteMenuByIds(c *gin.Context) {

	req := requests.DeleteMenuRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service menu.MenuService = &menu.MenuServiceImpl{}

	result := service.DeleteMenuById(req.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
