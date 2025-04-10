package system

import (
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type RoleDeptDao struct {
	db *gorm.DB
}

func NewRoleDeptDao(DB *gorm.DB) *RoleDeptDao {
	return &RoleDeptDao{
		db: DB,
	}
}

// CreateRoleDept 添加角色和部门关联
func (b RoleDeptDao) CreateRoleDept(dto system.AddRoleDeptDto) error {
	item := m.RoleDept{
		RoleId: dto.RoleId, // 角色id
		DeptId: dto.DeptId, // 部门id
	}

	return b.db.Create(&item).Error
}

// DeleteRoleDeptByIds 根据id删除角色和部门关联
func (b RoleDeptDao) DeleteRoleDeptByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.RoleDept{}).Error
}

// QueryRoleDeptDetail 查询角色和部门关联详情
func (b RoleDeptDao) QueryRoleDeptDetail(dto system.QueryRoleDeptDetailDto) (m.RoleDept, error) {
	var item m.RoleDept
	err := b.db.Where("id", dto.Id).First(&item).Error
	return item, err
}

// QueryRoleDeptList 查询角色和部门关联列表
func (b RoleDeptDao) QueryRoleDeptList(dto system.QueryRoleDeptListDto) ([]m.RoleDept, int64) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []m.RoleDept
	tx := b.db.Model(&m.RoleDept{})
	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	tx.Count(&total)
	return list, total
}
