package role

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
)

// RoleServiceImpl 角色信息操作实现
type RoleServiceImpl struct {
	Dao *system.RoleDao
}

func NewRoleServiceImpl(dao *system.RoleDao) RoleService {
	return &RoleServiceImpl{
		Dao: dao,
	}
}

// CreateRole 添加角色信息
func (s *RoleServiceImpl) CreateRole(dto d.AddRoleDto) error {
	return s.Dao.CreateRole(dto)
}

// DeleteRoleByIds 删除角色信息
func (s *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return s.Dao.DeleteRoleByIds(ids)
}

// UpdateRole 更新角色信息
func (s *RoleServiceImpl) UpdateRole(dto d.UpdateRoleDto) error {
	item, err := s.Dao.QueryRoleById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("角色信息不存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateRole(dto)
}

// UpdateRoleStatus 更新角色信息状态
func (s *RoleServiceImpl) UpdateRoleStatus(dto d.UpdateRoleStatusDto) error {
	return s.Dao.UpdateRoleStatus(dto)
}

// QueryRoleDetail 查询角色信息详情
func (s *RoleServiceImpl) QueryRoleDetail(dto d.QueryRoleDetailDto) (*d.QueryRoleListDtoResp, error) {
	item, err := s.Dao.QueryRoleDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("角色信息不存在")
	}

	return &d.QueryRoleListDtoResp{
		Id:         item.Id,                             // 主键
		RoleName:   item.RoleName,                       // 名称
		RoleKey:    item.RoleKey,                        // 角色权限字符串
		DataScope:  item.DataScope,                      // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     item.Status,                         // 状态(1:正常，0:禁用)
		Remark:     item.Remark,                         // 备注
		DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryRoleById 根据id查询角色信息详情
func (s *RoleServiceImpl) QueryRoleById(id int64) (*d.QueryRoleListDtoResp, error) {
	item, err := s.Dao.QueryRoleById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("角色信息不存在")
	}

	return &d.QueryRoleListDtoResp{
		Id:         item.Id,                             // 主键
		RoleName:   item.RoleName,                       // 名称
		RoleKey:    item.RoleKey,                        // 角色权限字符串
		DataScope:  item.DataScope,                      // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
		Status:     item.Status,                         // 状态(1:正常，0:禁用)
		Remark:     item.Remark,                         // 备注
		DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryRoleList 查询角色信息列表
func (s *RoleServiceImpl) QueryRoleList(dto d.QueryRoleListDto) ([]*d.QueryRoleListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryRoleList(dto)

	if err != nil {
		return nil, 0, err
	}

	var list []*d.QueryRoleListDtoResp

	for _, item := range result {
		resp := &d.QueryRoleListDtoResp{
			Id:         item.Id,                             // 主键
			RoleName:   item.RoleName,                       // 名称
			RoleKey:    item.RoleKey,                        // 角色权限字符串
			DataScope:  item.DataScope,                      // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			Remark:     item.Remark,                         // 备注
			DelFlag:    item.DelFlag,                        // 删除标志（0代表删除 1代表存在）
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
