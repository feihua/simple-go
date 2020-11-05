package menu

import (
	"simple-go/dto"
	"simple-go/pkg/util"
	"simple-go/repositories"
)

type MenuService struct {
}

func (menu *MenuService) CreateMenu(dto dto.MenuDto) error {
	return repositories.CreateSysMenu(dto)
}

func (menu *MenuService) GetMenuList(username string) []util.Tree {
	return repositories.GetSysMenuList()
}

func (menu *MenuService) UpdateMenu(menuDto dto.MenuDto) error {
	return repositories.UpdateSysMenu(menuDto)
}

func (menu *MenuService) DeletMenuById(id int) error {
	return repositories.DeleteSysMenuById(id)
}
