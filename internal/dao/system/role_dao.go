package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(DB *gorm.DB) *RoleDao {
	return &RoleDao{
		db: DB,
	}
}

// CreateRole 添加角色信息
func (b RoleDao) CreateRole(dto system.AddRoleDto) error {
	item := m.Role{
		RoleName:  dto.RoleName,  // 名称
		RoleKey:   dto.RoleKey,   // 角色权限字符串
		DataScope: dto.DataScope, // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:    dto.Status,    // 状态(1:正常，0:禁用)
		Remark:    dto.Remark,    // 备注
		DelFlag:   1,             // 删除标志（0代表删除 1代表存在）
		CreateBy:  dto.CreateBy,  // 创建者

	}

	return b.db.Create(&item).Error
}

// DeleteRoleByIds 根据id删除角色信息
func (b RoleDao) DeleteRoleByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.Role{}).Error
}

// UpdateRole 更新角色信息
func (b RoleDao) UpdateRole(dto system.UpdateRoleDto) error {

	item := m.Role{
		Id:         dto.Id,          // 主键
		RoleName:   dto.RoleName,    // 名称
		RoleKey:    dto.RoleKey,     // 角色权限字符串
		DataScope:  dto.DataScope,   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     dto.Status,      // 状态(1:正常，0:禁用)
		Remark:     dto.Remark,      // 备注
		DelFlag:    dto.DelFlag,     // 删除标志（0代表删除 1代表存在）
		CreateBy:   dto.CreateBy,    // 创建者
		CreateTime: dto.CreateTime,  // 创建时间
		UpdateBy:   dto.UpdateBy,    // 更新者
		UpdateTime: &dto.UpdateTime, // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateRoleStatus 更新角色信息状态
func (b RoleDao) UpdateRoleStatus(dto system.UpdateRoleStatusDto) error {

	return b.db.Model(&m.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryRoleDetail 查询角色信息详情
func (b RoleDao) QueryRoleDetail(dto system.QueryRoleDetailDto) (*m.Role, error) {
	var item m.Role
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryRoleList 查询角色信息列表
func (b RoleDao) QueryRoleList(dto system.QueryRoleListDto) ([]*m.Role, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.Role
	tx := b.db.Model(&m.Role{})
	if len(dto.RoleName) > 0 {
		tx.Where("role_name like %?%", dto.RoleName) // 名称
	}
	if len(dto.RoleKey) > 0 {
		tx.Where("role_key like %?%", dto.RoleKey) // 角色权限字符串
	}
	if dto.DataScope != 2 {
		tx.Where("data_scope=?", dto.DataScope) // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	}
	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 状态(1:正常，0:禁用)
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	err := tx.Count(&total).Error
	return list, total, err
}

// QueryRoleById 根据id查询角色
func (b RoleDao) QueryRoleById(id int64) (*m.Role, error) {
	var item m.Role
	err := b.db.Where("id = ?", id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryRoleByName 根据name查询角色
func (b RoleDao) QueryRoleByName(name string) (*m.Role, error) {
	var item m.Role
	err := b.db.Where("role_name = ?", name).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryRoleByKey 根据key查询角色
func (b RoleDao) QueryRoleByKey(key string) (*m.Role, error) {
	var item m.Role
	err := b.db.Where("role_key = ?", key).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
