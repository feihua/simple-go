package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type MenuDao struct {
	db *gorm.DB
}

func NewMenuDao(DB *gorm.DB) *MenuDao {
	return &MenuDao{db: DB}
}

// CreateMenu 创建菜单
func (m MenuDao) CreateMenu(dto system.MenuDto) error {
	menu := system2.Menu{
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
func (m MenuDao) QueryMenuList() ([]system2.Menu, error) {

	var menus []system2.Menu
	err := m.db.Model(&system2.Menu{}).Where("status_id = ?", 1).Find(&menus).Error

	return menus, err
}

// UpdateMenu 更新菜单
func (m MenuDao) UpdateMenu(dto system.MenuDto) error {
	menu := system2.Menu{
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
	err := m.db.Delete(&system2.Menu{Id: id}).Error
	return err
}
