package dao

import (
	"github.com/feihua/simple-go/models"
	"time"
)

func QueryRoleMenuList(roleId int64) []int64 {

	var roleMenus []models.RoleMenu
	models.DB.Where("role_id=?", roleId).Find(&roleMenus)

	var menuIds []int64

	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}

	return menuIds
}

func UpdateRoleMenu(roleId int64, menuIds []int64) {

	//先删除
	models.DB.Where("role_id = ?", roleId).Delete(&models.RoleMenu{})

	for _, menuId := range menuIds {

		roleMenu := models.RoleMenu{}
		roleMenu.RoleId = roleId
		roleMenu.MenuId = menuId
		roleMenu.CreateTime = time.Now()

		models.DB.Create(&roleMenu)

	}

}
