package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/model/system"
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

	var roleMenus []system.RoleMenu
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

	// 先删除
	err := r.db.Where("role_id = ?", roleId).Delete(&system.RoleMenu{}).Error
	if err != nil {
		return errors.New("删除角色关联菜单失败")
	}

	var list []system.RoleMenu
	for _, menuId := range menuIds {
		list = append(list, system.RoleMenu{
			RoleId:     roleId,
			MenuId:     menuId,
			CreateTime: time.Now(),
		})

	}
	return r.db.Model(&system.RoleMenu{}).CreateInBatches(list, len(list)).Error

}
