package repositories

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateDept(dto dto.DeptDto) error {
	dept := models.SysDept{}

	err := models.DB.Create(&dept).Error

	return err
}

func GetDeptList() ([]models.SysDept, int) {

	var total = 0
	var dept []models.SysDept
	models.DB.Find(&dept)

	models.DB.Model(&models.SysRole{}).Count(&total)
	return dept, total
}

func UpdateDept(deptDto dto.DeptDto) error {

	dept := models.SysDept{}
	if deptDto.Username != "" {
		dept.Name = deptDto.Username
	}

	err := models.DB.Model(&dept).Update(&dept).Error

	return err
}

func DeleteDeptById(id int) error {
	err := models.DB.Delete(&models.SysDept{Id: id}).Error
	return err
}
