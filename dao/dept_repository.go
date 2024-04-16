package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateDept(dto dto.DeptDto) error {
	dept := models.Dept{
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remarks:  dto.Remarks,
	}

	err := models.DB.Create(&dept).Error

	return err
}

func QueryDeptList() ([]models.Dept, int) {

	var total = 0
	var dept []models.Dept
	models.DB.Find(&dept)

	models.DB.Model(&models.Dept{}).Count(&total)
	return dept, total
}

func UpdateDept(dto dto.DeptDto) error {

	dept := models.Dept{
		Id:       dto.Id,
		DeptName: dto.DeptName,
		ParentId: dto.ParentId,
		Sort:     dto.Sort,
		Remarks:  dto.Remarks,
	}

	err := models.DB.Model(&dept).Update(&dept).Error

	return err
}

func DeleteDeptByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}

// QueryDeptByName 根据部门名称查询
func QueryDeptByName(name string) (*models.Dept, error) {

	var dept models.Dept
	err := models.DB.First(&dept).Where("dept_name = ?", name).Error

	return &dept, err
}
