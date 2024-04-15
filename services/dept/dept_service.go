package dept

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DeptService struct {
}

func (d *DeptService) CreateDept(dto dto.DeptDto) error {
	return dao.CreateDept(dto)
}

func (d *DeptService) QueryDeptList() ([]models.Dept, int) {
	return dao.QueryDeptList()
}

func (d *DeptService) UpdateDept(deptDto dto.DeptDto) error {
	return dao.UpdateDept(deptDto)
}

func (d *DeptService) DeleteDeptById(id int64) error {
	return dao.DeleteDeptById(id)
}
