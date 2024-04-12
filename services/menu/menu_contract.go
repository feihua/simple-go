package menu

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type MenuContract interface {
	CreateMenu(dto dto.MenuDto) error

	QueryMenuList(username string) ([]models.Menu, int)

	UpdateMenu(menuDto dto.MenuDto) error

	DeletMenuById(id int64) error
}
