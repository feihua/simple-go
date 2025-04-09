package dept

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
	"github.com/feihua/simple-go/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// DeptServiceImpl 部门操作实现
type DeptServiceImpl struct {
	deptDao *system.DeptDao
}

func NewDeptServiceImpl(dao *system.DeptDao) DeptService {
	return &DeptServiceImpl{
		deptDao: dao,
	}
}

// CreateDept 添加部门
func (s *DeptServiceImpl) CreateDept(dto a.AddDeptDto) error {
	return s.deptDao.CreateDept(dto)
}

// DeleteDeptByIds 删除部门
func (s *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return s.deptDao.DeleteDeptByIds(ids)
}

// UpdateDept 更新部门
func (s *DeptServiceImpl) UpdateDept(dto a.UpdateDeptDto) error {
	deptInfo, err := s.deptDao.QueryDeptById(dto.Id)

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		utils.Logger.Debugf("更新部门失败,部门不存在, 请求参数：%+v, 异常信息: %s", dto, err.Error())
		return errors.New("更新部门失败,部门不存在")
	case err != nil:
		utils.Logger.Debugf("查询部门异常, 请求参数：%+v, 异常信息: %s", dto, err.Error())
		return errors.New("查询部门异常")
	}

	dto.DelFlag = deptInfo.DelFlag
	dto.CreateBy = deptInfo.CreateBy
	dto.CreateTime = deptInfo.CreateTime
	dto.UpdateTime = time.Now()
	return s.deptDao.UpdateDept(dto)
}

// UpdateDeptStatus 更新部门状态
func (s *DeptServiceImpl) UpdateDeptStatus(dto a.UpdateDeptStatusDto) error {
	return s.deptDao.UpdateDeptStatus(dto)
}

// QueryDeptDetail 查询部门详情
func (s *DeptServiceImpl) QueryDeptDetail(dto a.QueryDeptDetailDto) (b.Dept, error) {
	return s.deptDao.QueryDeptDetail(dto)
}

// QueryDeptList 查询部门列表
func (s *DeptServiceImpl) QueryDeptList(dto a.QueryDeptListDto) ([]b.Dept, error) {
	return s.deptDao.QueryDeptList(dto)
}
