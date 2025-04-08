package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
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
	item := a.RoleMenu{
		Id:     dto.Id,     // 主键
		RoleId: dto.RoleId, // 角色ID
		MenuId: dto.MenuId, // 菜单ID
	}

	return b.db.Create(&item).Error
}

// DeleteRoleMenuByIds 根据id删除菜单角色关联
func (b RoleMenuDao) DeleteRoleMenuByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&a.RoleMenu{}).Error
}

// UpdateRoleMenu 更新菜单角色关联
func (b RoleMenuDao) UpdateRoleMenu(dto system.UpdateRoleMenuDto) error {

	item := a.RoleMenu{
		Id:     dto.Id,     // 主键
		RoleId: dto.RoleId, // 角色ID
		MenuId: dto.MenuId, // 菜单ID
	}

	return b.db.Updates(&item).Error
}

// UpdateRoleMenuStatus 更新菜单角色关联状态
func (b RoleMenuDao) UpdateRoleMenuStatus(dto system.UpdateRoleMenuStatusDto) error {

	return b.db.Model(&a.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryRoleMenuDetail 查询菜单角色关联详情
func (b RoleMenuDao) QueryRoleMenuDetail(dto system.QueryRoleMenuDetailDto) (a.RoleMenu, error) {
	var item a.RoleMenu
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryRoleMenuList 查询菜单角色关联列表
func (b RoleMenuDao) QueryRoleMenuList(dto system.QueryRoleMenuListDto) ([]a.RoleMenu, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []a.RoleMenu
	tx := b.db.Model(&a.RoleMenu{})
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
