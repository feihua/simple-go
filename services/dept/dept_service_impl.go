package dept

import (
	"errors"
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type DeptServiceImpl struct {
}

func (d *DeptServiceImpl) CreateDept(dto dto.DeptDto) error {
	//判断部门是否已经存在
	dept, err := dao.QueryDeptByName(dto.DeptName)
	if err != nil {
		return err
	}

	if dept != nil {
		return errors.New("部门已存在")
	}

	//添加部门
	return dao.CreateDept(dto)
}

func (d *DeptServiceImpl) QueryDeptList() ([]models.Dept, int) {
	return dao.QueryDeptList()
}

// QueryDeptByName 根据部门名称查询
func (d *DeptServiceImpl) QueryDeptByName(deptName string) (*models.Dept, error) {
	return dao.QueryDeptByName(deptName)
}

func (d *DeptServiceImpl) UpdateDept(deptDto dto.DeptDto) error {
	return dao.UpdateDept(deptDto)
}

func (d *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return dao.DeleteDeptByIds(ids)
}
