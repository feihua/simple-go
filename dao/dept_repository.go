package dao

import (
	"errors"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
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
func (d DeptDao) CreateDept(dto dto.DeptDto) error {
	dept := models.Dept{
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	return d.db.Create(&dept).Error
}

// QueryDeptList 查询部门列表(不需要分页)
func (d DeptDao) QueryDeptList() ([]models.Dept, error) {

	var dept []models.Dept
	err := d.db.Find(&dept).Error

	return dept, err
}

// UpdateDept 更新部门
func (d DeptDao) UpdateDept(dto dto.DeptDto) error {

	dept := models.Dept{
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
	return d.db.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}

// QueryDeptByName 根据部门名称查询
func (d DeptDao) QueryDeptByName(name string) (*models.Dept, error) {

	var dept models.Dept
	err := d.db.First(&dept).Where("dept_name = ?", name).Error

	switch {
	case err == nil:
		return &dept, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
