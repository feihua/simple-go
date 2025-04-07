package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	system2 "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type DeptDao struct {
	db *gorm.DB
}

func NewDeptDao(DB *gorm.DB) *DeptDao {
	return &DeptDao{
		db: DB,
	}
}

// CreateDept 创建部门
func (d DeptDao) CreateDept(dto system.DeptDto) error {
	dept := system2.Dept{
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	return d.db.Create(&dept).Error
}

// QueryDeptList 查询部门列表(不需要分页)
func (d DeptDao) QueryDeptList() ([]system2.Dept, error) {

	var dept []system2.Dept
	err := d.db.Find(&dept).Error

	return dept, err
}

// UpdateDept 更新部门
func (d DeptDao) UpdateDept(dto system.DeptDto) error {

	dept := system2.Dept{
		Id:       dto.Id,
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	return d.db.Model(&dept).Updates(&dept).Error
}

// DeleteDeptByIds 根据id删除部门
func (d DeptDao) DeleteDeptByIds(ids []int64) error {
	return d.db.Where("id in (?)", ids).Delete(&system2.Dept{}).Error
}

// QueryDeptByName 根据部门名称查询
func (d DeptDao) QueryDeptByName(name string) (*system2.Dept, error) {

	var dept system2.Dept
	err := d.db.First(&dept, "dept_name = ?", name).Error

	switch {
	case err == nil:
		return &dept, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
