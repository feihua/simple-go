package menu

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type MenuService struct {
}

func (menu *MenuService) CreateMenu(dto dto.MenuDto) error {
	return dao.CreateSysMenu(dto)
}

func (menu *MenuService) QueryMenuList(username string) ([]models.Menu, int) {
	return dao.QueryMenuList()
}

func (menu *MenuService) UpdateMenu(menuDto dto.MenuDto) error {
	return dao.UpdateSysMenu(menuDto)
}

func (menu *MenuService) DeletMenuById(id int64) error {
	return dao.DeleteSysMenuById(id)
}
