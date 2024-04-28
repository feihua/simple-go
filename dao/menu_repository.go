package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"gorm.io/gorm"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(DB *gorm.DB) *MenuDao {
	return &MenuDao{db: DB}
}

// CreateMenu 创建菜单
func (m MenuDao) CreateMenu(dto dto.MenuDto) error {
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

	return m.db.Create(&menu).Error
}

// QueryMenuList 查询菜单
func (m MenuDao) QueryMenuList() ([]models.Menu, int64) {

	var total int64 = 0
	var menus []models.Menu
	m.db.Find(&menus)

	m.db.Model(&models.Menu{}).Count(&total)
	return menus, total
}

// UpdateMenu 更新菜单
func (m MenuDao) UpdateMenu(dto dto.MenuDto) error {
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

	err := m.db.Model(&menu).Updates(&menu).Error

	return err
}

// DeleteMenuById 删除菜单
func (m MenuDao) DeleteMenuById(id int64) error {
	err := m.db.Delete(&models.Menu{Id: id}).Error
	return err
}
