package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateSysRole(dto dto.RoleDto) error {

	role := models.Role{
		RoleName: dto.RoleName,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := models.DB.Create(&role).Error

	return err
}

func QueryRoleList(current int, pageSize int) ([]models.Role, int) {

	var total = 0
	var role []models.Role
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&role)

	models.DB.Model(&models.Role{}).Count(&total)
	return role, total
}

// UpdateUser 更新用户信息
func UpdateSysRole(dto dto.RoleDto) error {

	role := models.Role{
		Id:       dto.Id,
		RoleName: dto.RoleName,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := models.DB.Model(&role).Update(&role).Error

	return err
}

// DeleteUserById 通过Id删除用户
func DeleteSysRoleById(id int64) error {
	err := models.DB.Delete(&models.Role{Id: id}).Error
	return err
}
