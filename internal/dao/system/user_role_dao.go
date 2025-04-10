package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
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
	item := m.UserRole{
		Id:     dto.Id,     // 主键
		UserId: dto.UserId, // 用户ID
		RoleId: dto.RoleId, // 角色ID
	}

	return b.db.Create(&item).Error
}

// DeleteUserRoleByIds 根据id删除角色用户关联
func (b UserRoleDao) DeleteUserRoleByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.UserRole{}).Error
}

// QueryUserRoleList 查询角色用户关联列表
func (b UserRoleDao) QueryUserRoleList(dto system.QueryUserRoleListDto) ([]m.UserRole, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []m.UserRole
	tx := b.db.Model(&m.UserRole{})
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

// IsAdministrator 根据用户id判断是否是管理员
func (u UserRoleDao) IsAdministrator(userId int64) bool {

	var count int64
	// 1是预留超级管理员角色的id
	u.db.Model(&m.UserRole{}).Where("user_id= ? and role_id = 1", userId).Count(&count)

	return count > 0
}
