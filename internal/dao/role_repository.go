package dao

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/models"
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
func (r RoleDao) QueryRoleList(roleDto dto.QueryRoleListDto) ([]models.Role, int64) {
	pageNo := roleDto.PageNo
	pageSize := roleDto.PageSize

	var total int64 = 0
	var role []models.Role
	tx := r.db.Model(&models.Role{})

	if roleDto.RoleName != "" {
		tx.Where("role_name=?", roleDto.RoleName)
	}
	if roleDto.StatusId != 2 {
		tx.Where("status_id=?", roleDto.StatusId)
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&role)

	tx.Count(&total)
	return role, total
}

// QueryAllRoleList 查询角色列表
func (r RoleDao) QueryAllRoleList() ([]models.Role, error) {

	var role []models.Role
	err := r.db.Model(&models.Role{}).Where("status_id=?", 1).Find(&role).Error

	return role, err
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
	switch {
	case err == nil:
		return &role, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
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
