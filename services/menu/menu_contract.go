package menu

import (
	"simple-go/dto"
	"simple-go/pkg/util"
)

type MenuContract interface {

	CreateMenu(dto dto.MenuDto) error

	GetMenuList(username string) []util.Tree

	UpdateMenu(menuDto dto.MenuDto) error

	DeletMenuById(id int) error
}
