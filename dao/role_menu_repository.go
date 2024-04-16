package dao

import (
	"github.com/feihua/simple-go/models"
	"github.com/jinzhu/gorm"
	"time"
)

type RoleMenuDao struct {
	db *gorm.DB
}

func NewRoleMenuDao(DB *gorm.DB) *RoleMenuDao {
	return &RoleMenuDao{db: DB}
}

// QueryRoleMenuList 查询角色菜单
func (r RoleMenuDao) QueryRoleMenuList(roleId int64) []int64 {

	var roleMenus []models.RoleMenu
	r.db.Where("role_id=?", roleId).Find(&roleMenus)

	var menuIds []int64

	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}

	return menuIds
}

// UpdateRoleMenuList 更新角色菜单
func (r RoleMenuDao) UpdateRoleMenuList(roleId int64, menuIds []int64) {

	//先删除
	r.db.Where("role_id = ?", roleId).Delete(&models.RoleMenu{})

	for _, menuId := range menuIds {

		roleMenu := models.RoleMenu{}
		roleMenu.RoleId = roleId
		roleMenu.MenuId = menuId
		roleMenu.CreateTime = time.Now()

		r.db.Create(&roleMenu)

	}

}
