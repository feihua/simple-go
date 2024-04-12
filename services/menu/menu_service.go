package menu

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
)

type MenuService struct {
}

func (menu *MenuService) CreateMenu(dto dto.MenuDto) error {
	return repositories.CreateSysMenu(dto)
}

func (menu *MenuService) GetMenuList(username string) ([]models.SysMenu, int) {
	return repositories.GetSysMenuList()
}

func (menu *MenuService) UpdateMenu(menuDto dto.MenuDto) error {
	return repositories.UpdateSysMenu(menuDto)
}

func (menu *MenuService) DeletMenuById(id int) error {
	return repositories.DeleteSysMenuById(id)
}
