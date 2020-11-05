package repositories

import (
	"simple-go/dto"
	"simple-go/models"
)

func CreateSysRole(dto dto.RoleDto) error {
	user := models.SysRole{}

	err := models.DB.Create(&user).Error

	return err
}

func GetSysRoleList(current int, pageSize int) ([]models.SysRole, int) {

	var total = 0
	var role []models.SysRole
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&role)

	models.DB.Model(&models.SysRole{}).Count(&total)
	return role, total
}

//UpdateUser 更新用户信息
func UpdatSysRole(roleDto dto.RoleDto) error {
	role := models.SysRole{}

	err := models.DB.Model(&role).Update(&role).Error

	return err
}

//DeleteUserById 通过Id删除用户
func DeleteSysRoleById(id int64) error {
	err := models.DB.Delete(&models.SysRole{Id: id}).Error
	return err
}
