package dept

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// DeptServiceImpl 部门操作实现
type DeptServiceImpl struct {
	Dao *system.DeptDao
}

func NewDeptServiceImpl(dao *system.DeptDao) DeptService {
	return &DeptServiceImpl{
		Dao: dao,
	}
}

// CreateDept 添加部门
func (s *DeptServiceImpl) CreateDept(dto a.AddDeptDto) error {
	return s.Dao.CreateDept(dto)
}

// DeleteDeptByIds 删除部门
func (s *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return s.Dao.DeleteDeptByIds(ids)
}

// UpdateDept 更新部门
func (s *DeptServiceImpl) UpdateDept(dto a.UpdateDeptDto) error {
	return s.Dao.UpdateDept(dto)
}

// UpdateDeptStatus 更新部门状态
func (s *DeptServiceImpl) UpdateDeptStatus(dto a.UpdateDeptStatusDto) error {
	return s.Dao.UpdateDeptStatus(dto)
}

// QueryDeptDetail 查询部门详情
func (s *DeptServiceImpl) QueryDeptDetail(dto a.QueryDeptDetailDto) (b.Dept, error) {
	return s.Dao.QueryDeptDetail(dto)
}

// QueryDeptList 查询部门列表
func (s *DeptServiceImpl) QueryDeptList(dto a.QueryDeptListDto) ([]b.Dept, int64) {
	return s.Dao.QueryDeptList(dto)
}
