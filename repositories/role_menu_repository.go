package repositories

import (
	"github.com/feihua/simple-go/models"
	"time"
)

func GeRoleMenuList(roleId int64) []int64 {

	var roleMenus []models.SysRoleMenu
	models.DB.Where("role_id=?", roleId).Find(&roleMenus)

	var menuIds []int64

	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}

	return menuIds
}

func UpdateRoleMenu(roleId int64, menuIds []int64) {

	//先删除
	models.DB.Raw("DELETE FROM sys_role_menu WHERE role_id = ?", roleId)

	for _, menuId := range menuIds {

		roleMenu := models.SysRoleMenu{}
		roleMenu.RoleId = roleId
		roleMenu.MenuId = menuId
		roleMenu.CreateBy = "admin"
		roleMenu.CreateTime = time.Now()

		models.DB.Create(&roleMenu)

	}

}
