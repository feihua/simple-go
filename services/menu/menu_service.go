package menu

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// MenuService 菜单操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:42
*/
type MenuService interface {
	// CreateMenu 创建菜单
	CreateMenu(dto dto.MenuDto) error

	// QueryMenuList 查询菜单
	QueryMenuList() ([]models.Menu, error)

	// UpdateMenu 更新菜单
	UpdateMenu(menuDto dto.MenuDto) error

	// DeleteMenuById 删除菜单
	DeleteMenuById(id int64) error
}
