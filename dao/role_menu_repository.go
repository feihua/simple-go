package dao

import (
	"errors"
	"github.com/feihua/simple-go/models"
	"gorm.io/gorm"
	"time"
)

type RoleMenuDao struct {
	db *gorm.DB
}

func NewRoleMenuDao(DB *gorm.DB) *RoleMenuDao {
	return &RoleMenuDao{db: DB}
}

// QueryRoleMenuList 查询角色菜单
func (r RoleMenuDao) QueryRoleMenuList(roleId int64) ([]int64, error) {

	var roleMenus []models.RoleMenu
	err := r.db.Where("role_id=?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, errors.New("查询角色菜单失败")
	}

	var menuIds []int64

	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}

	return menuIds, nil
}

// UpdateRoleMenuList 更新角色菜单
func (r RoleMenuDao) UpdateRoleMenuList(roleId int64, menuIds []int64) error {

	//先删除
	err := r.db.Where("role_id = ?", roleId).Delete(&models.RoleMenu{}).Error
	if err != nil {
		return errors.New("删除角色关联菜单失败")
	}

	var list []models.RoleMenu
	for _, menuId := range menuIds {
		list = append(list, models.RoleMenu{
			RoleId:     roleId,
			MenuId:     menuId,
			CreateTime: time.Now(),
		})

	}
	return r.db.Model(&models.RoleMenu{}).CreateInBatches(list, len(list)).Error

}
