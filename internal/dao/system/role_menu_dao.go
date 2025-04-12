package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type RoleMenuDao struct {
	db *gorm.DB
}

func NewRoleMenuDao(DB *gorm.DB) *RoleMenuDao {
	return &RoleMenuDao{
		db: DB,
	}
}

// CreateRoleMenu 添加菜单角色关联
func (b RoleMenuDao) CreateRoleMenu(dto system.AddRoleMenuDto) error {
	item := m.RoleMenu{
		RoleId: dto.RoleId, // 角色ID
		MenuId: dto.MenuId, // 菜单ID
	}

	return b.db.Create(&item).Error
}

// CreateRoleMenuBatch 批量添加角色用户关联
func (b RoleMenuDao) CreateRoleMenuBatch(dto system.UpdateRoleMenuDto) error {
	var userRoles []*m.RoleMenu
	for _, menuId := range dto.MenuIds {
		sysUserRole := m.RoleMenu{
			RoleId: dto.RoleId,
			MenuId: menuId,
		}
		userRoles = append(userRoles, &sysUserRole)
	}
	return b.db.CreateInBatches(&userRoles, len(userRoles)).Error
}

// DeleteRoleMenuByIds 根据id删除菜单角色关联
func (b RoleMenuDao) DeleteRoleMenuByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.RoleMenu{}).Error
}

// QueryRoleMenuDetail 查询菜单角色关联详情
func (b RoleMenuDao) QueryRoleMenuDetail(dto system.QueryRoleMenuDetailDto) (m.RoleMenu, error) {
	var item m.RoleMenu
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// DeleteRoleMenuByRoleId 根据roleId删除角色用户关联
func (b RoleMenuDao) DeleteRoleMenuByRoleId(roleId int64) error {
	return b.db.Where("role_id = ?", roleId).Delete(&m.RoleMenu{}).Error
}

// QueryMenuIds 根据roleId查询菜单ids
func (b RoleMenuDao) QueryMenuIds(roleId int64) ([]int64, error) {
	var ids []int64
	err := b.db.Model(&m.RoleMenu{}).Select("menu_id").Where("role_id", roleId).Scan(&ids).Error
	return ids, err
}
