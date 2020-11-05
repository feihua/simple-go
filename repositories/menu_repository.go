package repositories

import (
	"simple-go/dto"
	"simple-go/models"
	"simple-go/pkg/util"
)


func CreateSysMenu(dto dto.MenuDto) error {
	menu := models.SysMenu{}

	err := models.DB.Create(&menu).Error

	return err
}

func GetSysMenuList() []util.Tree {

	var menu []models.SysMenu
	models.DB.Find(&menu)

	// 生成完全树
	resp := util.GenerateTree(models.SystemMenus.ConvertToINodeArray(menu), nil)

	return resp
}

func UpdateSysMenu(menuDto dto.MenuDto) error {
	menu := models.SysMenu{}

	err := models.DB.Model(&menu).Update(&menu).Error

	return err
}

func DeleteSysMenuById(id int) error {
	err := models.DB.Delete(&models.SysMenu{Id: id}).Error
	return err
}
