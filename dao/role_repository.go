package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"gorm.io/gorm"
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
func (r RoleDao) QueryRoleList(current int, pageSize int) ([]models.Role, int64) {

	var total int64 = 0
	var role []models.Role
	r.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&role)

	r.db.Model(&models.Role{}).Count(&total)
	return role, total
}

// QueryRoleById 根据角色Id查询角色
func (r RoleDao) QueryRoleById(roleId int64) (*models.Role, error) {

	var role models.Role

	err := r.db.First(&role, r.db.Where("id = ?", roleId)).Error
	return &role, err
}

// QueryRoleByName 根据角色名称查询角色
func (r RoleDao) QueryRoleByName(name string) (*models.Role, error) {

	var role models.Role

	err := r.db.First(&role, r.db.Where("role_name = ?", name)).Error
	return &role, err
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

	return r.db.Model(&role).Updates(&role).Error
}

// DeleteRoleByIds 通过Id删除角色
func (r RoleDao) DeleteRoleByIds(ids []int64) error {
	return r.db.Where("id in (?)", ids).Delete(&models.Role{}).Error
}
