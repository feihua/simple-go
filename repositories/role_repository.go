package repositories

import (
	"simple-go/dto"
	"simple-go/models"
	"time"
)

func CreateSysRole(dto dto.RoleDto) error {

	role := models.SysRole{}
	role.Name = dto.Name
	role.Remark = dto.Remark
	role.CreateBy = "liufeihua"
	role.CreateTime = time.Now()
	role.LastUpdateBy = "liufeihua"
	role.LastUpdateTime = time.Now()

	err := models.DB.Create(&role).Error

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
func UpdatSysRole(dto dto.RoleDto) error {

	role := models.SysRole{}
	role.Id = dto.Id
	role.Name = dto.Name
	role.Remark = dto.Remark
	role.CreateBy = "liufeihua"
	role.CreateTime = time.Now()
	role.LastUpdateBy = "liufeihua"
	role.LastUpdateTime = time.Now()

	err := models.DB.Model(&role).Update(&role).Error

	return err
}

//DeleteUserById 通过Id删除用户
func DeleteSysRoleById(id int64) error {
	err := models.DB.Delete(&models.SysRole{Id: id}).Error
	return err
}
