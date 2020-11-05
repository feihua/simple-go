package repositories

import (
	"simple-go/dto"
	"simple-go/models"
	"simple-go/pkg/util"
)

func CreateDept(dto dto.DeptDto) error {
	dept := models.SysDept{}


	err := models.DB.Create(&dept).Error

	return err
}

func GetDeptList() []util.Tree {

	var dept []models.SysDept
	models.DB.Find(&dept)

	// 生成完全树
	resp := util.GenerateTree(models.SysDepts.ConvertToINodeArray(dept), nil)
	return resp
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
