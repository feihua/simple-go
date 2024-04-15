package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateSysMenu(dto dto.MenuDto) error {
	menu := models.Menu{
		MenuName: dto.MenuName,
		MenuType: dto.MenuType,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		ParentId: dto.ParentId,
		MenuUrl:  dto.MenuUrl,
		ApiUrl:   dto.ApiUrl,
		MenuIcon: dto.MenuIcon,
		Remark:   dto.Remark,
	}

	err := models.DB.Create(&menu).Error

	return err
}

func QueryMenuList() ([]models.Menu, int) {

	var total = 0
	var menus []models.Menu
	models.DB.Find(&menus)

	models.DB.Model(&models.Menu{}).Count(&total)
	return menus, total
}

func UpdateSysMenu(dto dto.MenuDto) error {
	menu := models.Menu{
		Id:       dto.Id,
		MenuName: dto.MenuName,
		MenuType: dto.MenuType,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		ParentId: dto.ParentId,
		MenuUrl:  dto.MenuUrl,
		ApiUrl:   dto.ApiUrl,
		MenuIcon: dto.MenuIcon,
		Remark:   dto.Remark,
	}

	err := models.DB.Model(&menu).Update(&menu).Error

	return err
}

func DeleteSysMenuByIds(id int64) error {
	err := models.DB.Delete(&models.Menu{Id: id}).Error
	return err
}
