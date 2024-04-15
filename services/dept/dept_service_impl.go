package dept

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DeptServiceImpl struct {
}

func (d *DeptServiceImpl) CreateDept(dto dto.DeptDto) error {
	return dao.CreateDept(dto)
}

func (d *DeptServiceImpl) QueryDeptList() ([]models.Dept, int) {
	return dao.QueryDeptList()
}

func (d *DeptServiceImpl) UpdateDept(deptDto dto.DeptDto) error {
	return dao.UpdateDept(deptDto)
}

func (d *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return dao.DeleteDeptByIds(ids)
}
