package repositories

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateSysMenu(dto dto.MenuDto) error {
	menu := models.SysMenu{}

	err := models.DB.Create(&menu).Error

	return err
}

func GetSysMenuList() ([]models.SysMenu, int) {

	var total = 0
	var menus []models.SysMenu
	models.DB.Find(&menus)

	models.DB.Model(&models.SysMenu{}).Count(&total)
	return menus, total
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
