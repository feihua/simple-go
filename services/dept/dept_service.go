package dept

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
)

type DeptService struct {
}

func (d *DeptService) CreateDept(dto dto.DeptDto) error {
	return repositories.CreateDept(dto)
}

func (d *DeptService) QueryDeptList() ([]models.Dept, int) {
	return repositories.QueryDeptList()
}

func (d *DeptService) UpdateDept(deptDto dto.DeptDto) error {
	return repositories.UpdateDept(deptDto)
}

func (d *DeptService) DeleteDeptById(id int64) error {
	return repositories.DeleteDeptById(id)
}
