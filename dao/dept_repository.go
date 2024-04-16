package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// CreateDept 创建部门
func CreateDept(dto dto.DeptDto) error {
	dept := models.Dept{
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remarks:  dto.Remarks,
	}

	return models.DB.Create(&dept).Error
}

// QueryDeptList 查询部门列表(不需要分页)
func QueryDeptList() ([]models.Dept, error) {

	var dept []models.Dept
	err := models.DB.Find(&dept).Error

	return dept, err
}

// UpdateDept 更新部门
func UpdateDept(dto dto.DeptDto) error {

	dept := models.Dept{
		Id:       dto.Id,
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remarks:  dto.Remarks,
	}

	return models.DB.Model(&dept).Update(&dept).Error
}

// DeleteDeptByIds 根据id删除部门
func DeleteDeptByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}

// QueryDeptByName 根据部门名称查询
func QueryDeptByName(name string) (*models.Dept, error) {

	var dept models.Dept
	err := models.DB.First(&dept).Where("dept_name = ?", name).Error

	return &dept, err
}
