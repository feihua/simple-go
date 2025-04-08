package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type UserRoleDao struct {
	db *gorm.DB
}

func NewUserRoleDao(DB *gorm.DB) *UserRoleDao {
	return &UserRoleDao{
		db: DB,
	}
}

// CreateUserRole 添加角色用户关联
func (b UserRoleDao) CreateUserRole(dto system.AddUserRoleDto) error {
	item := a.UserRole{
		Id:     dto.Id,     // 主键
		UserId: dto.UserId, // 用户ID
		RoleId: dto.RoleId, // 角色ID
	}

	return b.db.Create(&item).Error
}

// DeleteUserRoleByIds 根据id删除角色用户关联
func (b UserRoleDao) DeleteUserRoleByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.UserRole{}).Error
}

// UpdateUserRole 更新角色用户关联
func (b UserRoleDao) UpdateUserRole(dto system.UpdateUserRoleDto) error {

	item := a.UserRole{
		Id:     dto.Id,     // 主键
		UserId: dto.UserId, // 用户ID
		RoleId: dto.RoleId, // 角色ID
	}

	return b.db.Updates(&item).Error
}

// UpdateUserRoleStatus 更新角色用户关联状态
func (b UserRoleDao) UpdateUserRoleStatus(dto system.UpdateUserRoleStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryUserRoleDetail 查询角色用户关联详情
func (b UserRoleDao) QueryUserRoleDetail(dto system.QueryUserRoleDetailDto) (a.UserRole, error) {
	var item a.UserRole
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryUserRoleList 查询角色用户关联列表
func (b UserRoleDao) QueryUserRoleList(dto system.QueryUserRoleListDto) ([]a.UserRole, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.UserRole
	tx := b.db.Model(&a.UserRole{})
	if dto.UserId != 2 {
		tx.Where("user_id=?", dto.UserId) // 用户ID
	}
	if dto.RoleId != 2 {
		tx.Where("role_id=?", dto.RoleId) // 角色ID
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
