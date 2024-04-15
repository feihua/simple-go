package menu

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type MenuServiceImpl struct {
}

func (menu *MenuServiceImpl) CreateMenu(dto dto.MenuDto) error {
	return dao.CreateSysMenu(dto)
}

func (menu *MenuServiceImpl) QueryMenuList(username string) ([]models.Menu, int) {
	return dao.QueryMenuList()
}

func (menu *MenuServiceImpl) UpdateMenu(menuDto dto.MenuDto) error {
	return dao.UpdateSysMenu(menuDto)
}

func (menu *MenuServiceImpl) DeleteMenuById(id int64) error {
	return dao.DeleteSysMenuByIds(id)
}
