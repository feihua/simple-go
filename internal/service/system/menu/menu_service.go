package menu

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
)

// MenuService 菜单操作接口
/*
Author: LiuFeiHua
Date: 2024/4/16 13:42
*/
type MenuService interface {
	// CreateMenu 创建菜单
	CreateMenu(dto system.MenuDto) error

	// QueryMenuList 查询菜单
	QueryMenuList() ([]system2.Menu, error)

	// UpdateMenu 更新菜单
	UpdateMenu(menuDto system.MenuDto) error

	// DeleteMenuById 删除菜单
	DeleteMenuById(id int64) error
}
