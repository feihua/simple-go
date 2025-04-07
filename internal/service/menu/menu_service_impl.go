package menu

import (
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
)

// MenuServiceImpl 菜单操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 13:41
*/
type MenuServiceImpl struct {
	Dao *dao.MenuDao
}

func NewMenuServiceImpl(Dao *dao.MenuDao) MenuService {
	return &MenuServiceImpl{Dao: Dao}
}

// CreateMenu 创建菜单
func (menu *MenuServiceImpl) CreateMenu(dto dto.MenuDto) error {
	return menu.Dao.CreateMenu(dto)
}

// QueryMenuList 查询菜单
func (menu *MenuServiceImpl) QueryMenuList() ([]model.Menu, error) {
	return menu.Dao.QueryMenuList()
}

// UpdateMenu 更新菜单
func (menu *MenuServiceImpl) UpdateMenu(menuDto dto.MenuDto) error {
	return menu.Dao.UpdateMenu(menuDto)
}

// DeleteMenuById 删除菜单
func (menu *MenuServiceImpl) DeleteMenuById(id int64) error {
	return menu.Dao.DeleteMenuById(id)
}
