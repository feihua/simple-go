package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/jinzhu/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(DB *gorm.DB) *RoleDao {
	return &RoleDao{db: DB}
}

// CreateRole 创建角色
func (r RoleDao) CreateRole(dto dto.RoleDto) error {

	role := models.Role{
		RoleName: dto.RoleName,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := r.db.Create(&role).Error

	return err
}

// QueryRoleList 查询角色列表
func (r RoleDao) QueryRoleList(current int, pageSize int) ([]models.Role, int) {

	var total = 0
	var role []models.Role
	r.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&role)

	r.db.Model(&models.Role{}).Count(&total)
	return role, total
}

// UpdateRole 更新角色
func (r RoleDao) UpdateRole(dto dto.RoleDto) error {

	role := models.Role{
		Id:       dto.Id,
		RoleName: dto.RoleName,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := r.db.Model(&role).Update(&role).Error

	return err
}

// DeleteRoleByIds 通过Id删除角色
func (r RoleDao) DeleteRoleByIds(ids []int64) error {
	return r.db.Where("id in (?)", ids).Delete(&models.Role{}).Error
}
