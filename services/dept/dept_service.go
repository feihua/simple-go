package dept

import (
	"simple-go/dto"
	"simple-go/pkg/util"
	"simple-go/repositories"
)

type DeptService struct {
}

func (d *DeptService) CreateDept(dto dto.DeptDto) error {
	return repositories.CreateDept(dto)
}

func (d *DeptService) GetDeptList() []util.Tree {
	return repositories.GetDeptList()
}

func (d *DeptService) UpdateDept(deptDto dto.DeptDto) error {
	return repositories.UpdateDept(deptDto)
}

func (d *DeptService) DeleteDeptById(id int) error {
	return repositories.DeleteDeptById(id)
}
