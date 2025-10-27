package dept

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
)

// DeptServiceImpl 部门操作实现
type DeptServiceImpl struct {
	Dao     *system.DeptDao
	UserDao *system.UserDao
}

func NewDeptServiceImpl(dao *system.DeptDao, UserDao *system.UserDao) DeptService {
	return &DeptServiceImpl{
		Dao:     dao,
		UserDao: UserDao,
	}
}

// CreateDept 添加部门
func (s *DeptServiceImpl) CreateDept(dto d.AddDeptDto) error {
	byName, err := s.Dao.QueryDeptByNameAndParentId(dto.DeptName, dto.ParentId)
	if err != nil {
		return err
	}
	if byName != nil {
		return errors.New("部门名称已存在")
	}

	parentDept, err := s.Dao.QueryDeptById(dto.ParentId)
	if err != nil {
		return err
	}

	if parentDept == nil {
		return errors.New("上级部门不存在")
	}

	if parentDept.Status != 1 {
		return errors.New("部门停用,不允许新增")
	}

	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.Id)
	dto.Ancestors = ancestors
	return s.Dao.CreateDept(dto)
}

// DeleteDeptById 删除部门
func (s *DeptServiceImpl) DeleteDeptById(id int64) error {
	if id == 1 {
		return errors.New("顶级部门,不允许删除")
	}

	item, err := s.Dao.QueryDeptById(id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("部门不存在")
	}

	if item.Status == 1 {
		return errors.New("部门状态为启用,不允许删除")
	}

	childDept, err := s.Dao.QueryChildDept(id)
	if err != nil {
		return err
	}

	if len(childDept) > 0 {
		return errors.New("存在下级部门,不允许删除")
	}

	users, err := s.UserDao.ExistUserByDeptId(item.Id)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.New("部门存在用户,不允许删除")
	}

	return s.Dao.DeleteDeptById(id)
}

// UpdateDept 更新部门
func (s *DeptServiceImpl) UpdateDept(dto d.UpdateDeptDto) error {
	oldDept, err := s.Dao.QueryDeptById(dto.Id)

	if err != nil {
		return err
	}

	if oldDept == nil {
		return errors.New("部门不存在")
	}

	if dto.ParentId == dto.Id {
		return errors.New("上级部门不能为当前部门")
	}

	parentDept, err := s.Dao.QueryDeptById(dto.ParentId)

	if err != nil {
		return err
	}

	if parentDept == nil {
		return errors.New("上级部门不存在")
	}

	// 3.根据部门名称查询部门是否已存在
	dept, err := s.Dao.QueryDeptByNameAndParentId(dto.DeptName, dto.ParentId)
	if err != nil {
		return err
	}
	if dept != nil && dept.Id != dto.Id {
		return errors.New("部门名称已存在")
	}

	count, err := s.Dao.QueryChildDeptCount(dto.Id)
	if err != nil {
		return err
	}
	if count > 0 && dto.Status == 0 {
		return errors.New("该部门包含未停用的子部门")
	}

	deptList, err := s.Dao.QueryChildDeptList(dto.Id)
	if err != nil {
		return err
	}

	// 5.修改下级部门祖级
	ancestors := fmt.Sprintf("%s,%d", parentDept.Ancestors, parentDept.Id)
	if len(deptList) > 0 {
		for _, dept1 := range deptList {
			parentIdStr := strings.Replace(dept1.Ancestors, oldDept.Ancestors, ancestors, -1)
			err = s.Dao.UpdateDeptAncestors(dept1.Id, parentIdStr)
			if err != nil {
				return err
			}
		}
	}

	if dto.Status == 1 {
		split := strings.Split(ancestors, ",")
		var parentIds []int64

		for _, str := range split {
			num, _ := strconv.ParseInt(str, 10, 64)
			parentIds = append(parentIds, num)
		}

		err = s.Dao.UpdateDeptStatus(parentIds, dto.Status, dto.UpdateBy)
		if err != nil {
			return err
		}
	}

	dto.Ancestors = ancestors
	dto.CreateBy = oldDept.CreateBy
	dto.CreateTime = oldDept.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateDept(dto)
}

// UpdateDeptStatus 更新部门状态
func (s *DeptServiceImpl) UpdateDeptStatus(dto d.UpdateDeptStatusDto) error {
	item, err := s.Dao.QueryDeptById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("部门不存在")
	}

	count, err := s.Dao.QueryChildDeptCount(dto.Id)
	if err != nil {
		return err
	}

	if count > 0 && dto.Status == 0 {
		return errors.New("该部门包含未停用的子部门")
	}

	var parentIds []int64
	if dto.Status == 1 {
		split := strings.Split(item.Ancestors, ",")
		for _, str := range split {
			num, _ := strconv.ParseInt(str, 10, 64)
			parentIds = append(parentIds, num)
		}

	}

	parentIds = append(parentIds, dto.Id)
	return s.Dao.UpdateDeptStatus(parentIds, dto.Status, dto.UpdateBy)

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
			Key:        strconv.FormatInt(item.Id, 10),      // 部门id
			ParentId:   item.ParentId,                       // 父部门id
			Ancestors:  item.Ancestors,                      // 祖级列表
			DeptName:   item.DeptName,                       // 部门名称
			Title:      item.DeptName,                       // 部门名称
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
