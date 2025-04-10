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
		Id:     dto.Id,     // 主键
		RoleId: dto.RoleId, // 角色ID
		MenuId: dto.MenuId, // 菜单ID
	}

	return b.db.Create(&item).Error
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

// QueryRoleMenuList 查询菜单角色关联列表
func (b RoleMenuDao) QueryRoleMenuList(dto system.QueryRoleMenuListDto) ([]m.RoleMenu, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []m.RoleMenu
	tx := b.db.Model(&m.RoleMenu{})
	if dto.RoleId != 2 {
		tx.Where("role_id=?", dto.RoleId) // 角色ID
	}
	if dto.MenuId != 2 {
		tx.Where("menu_id=?", dto.MenuId) // 菜单ID
	}
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
