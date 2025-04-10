package dept

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
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
func (s *DeptServiceImpl) CreateDept(dto d.AddDeptDto) error {
	return s.Dao.CreateDept(dto)
}

// DeleteDeptByIds 删除部门
func (s *DeptServiceImpl) DeleteDeptByIds(ids []int64) error {
	return s.Dao.DeleteDeptByIds(ids)
}

// UpdateDept 更新部门
func (s *DeptServiceImpl) UpdateDept(dto d.UpdateDeptDto) error {
	item, err := s.Dao.QueryDeptById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("部门不存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateDept(dto)
}

// UpdateDeptStatus 更新部门状态
func (s *DeptServiceImpl) UpdateDeptStatus(dto d.UpdateDeptStatusDto) error {
	return s.Dao.UpdateDeptStatus(dto)
}

// QueryDeptDetail 查询部门详情
func (s *DeptServiceImpl) QueryDeptDetail(dto d.QueryDeptDetailDto) (*d.QueryDeptListDtoResp, error) {
	item, err := s.Dao.QueryDeptDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("部门不存在")
	}

	return &d.QueryDeptListDtoResp{
		Id:         item.Id,                             // 部门id
		ParentId:   item.ParentId,                       // 父部门id
		Ancestors:  item.Ancestors,                      // 祖级列表
		DeptName:   item.DeptName,                       // 部门名称
		Sort:       item.Sort,                           // 显示顺序
		Leader:     item.Leader,                         // 负责人
		Phone:      item.Phone,                          // 联系电话
		Email:      item.Email,                          // 邮箱
		Status:     item.Status,                         // 部门状态（0：停用，1:正常）
		DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDeptById 根据id查询部门详情
func (s *DeptServiceImpl) QueryDeptById(id int64) (*d.QueryDeptListDtoResp, error) {
	item, err := s.Dao.QueryDeptById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("部门不存在")
	}

	return &d.QueryDeptListDtoResp{
		Id:         item.Id,                             // 部门id
		ParentId:   item.ParentId,                       // 父部门id
		Ancestors:  item.Ancestors,                      // 祖级列表
		DeptName:   item.DeptName,                       // 部门名称
		Sort:       item.Sort,                           // 显示顺序
		Leader:     item.Leader,                         // 负责人
		Phone:      item.Phone,                          // 联系电话
		Email:      item.Email,                          // 邮箱
		Status:     item.Status,                         // 部门状态（0：停用，1:正常）
		DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryDeptList 查询部门列表
func (s *DeptServiceImpl) QueryDeptList(dto d.QueryDeptListDto) ([]*d.QueryDeptListDtoResp, error) {
	result, err := s.Dao.QueryDeptList(dto)

	if err != nil {
		return nil, err
	}

	var list []*d.QueryDeptListDtoResp

	for _, item := range result {
		resp := &d.QueryDeptListDtoResp{
			Id:         item.Id,                             // 部门id
			ParentId:   item.ParentId,                       // 父部门id
			Ancestors:  item.Ancestors,                      // 祖级列表
			DeptName:   item.DeptName,                       // 部门名称
			Sort:       item.Sort,                           // 显示顺序
			Leader:     item.Leader,                         // 负责人
			Phone:      item.Phone,                          // 联系电话
			Email:      item.Email,                          // 邮箱
			Status:     item.Status,                         // 部门状态（0：停用，1:正常）
			DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, nil
}
