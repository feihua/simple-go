package menu

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
)

// MenuServiceImpl 菜单信息操作实现
type MenuServiceImpl struct {
	Dao *system.MenuDao
}

func NewMenuServiceImpl(dao *system.MenuDao) MenuService {
	return &MenuServiceImpl{
		Dao: dao,
	}
}

// CreateMenu 添加菜单信息
func (s *MenuServiceImpl) CreateMenu(dto d.AddMenuDto) error {
	byName, err := s.Dao.QueryMenuByName(dto.MenuName)
	if err != nil {
		return err
	}
	if byName != nil {
		return errors.New("菜单名称已存在")
	}

	if len(dto.MenuUrl) != 0 {
		byUrl, err1 := s.Dao.QueryMenuByMenuUrl(dto.MenuUrl)
		if err1 != nil {
			return err1
		}
		if byUrl != nil {
			return errors.New("菜单URL已存在")
		}
	}

	return s.Dao.CreateMenu(dto)
}

// DeleteMenuByIds 删除菜单信息
func (s *MenuServiceImpl) DeleteMenuByIds(ids []int64) error {
	return s.Dao.DeleteMenuByIds(ids)
}

// UpdateMenu 更新菜单信息
func (s *MenuServiceImpl) UpdateMenu(dto d.UpdateMenuDto) error {
	item, err := s.Dao.QueryMenuById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("菜单信息不存在")
	}

	byName, err := s.Dao.QueryMenuByName(dto.MenuName)
	if err != nil {
		return err
	}
	if byName != nil && item.Id != dto.Id {
		return errors.New("菜单名称已存在")
	}

	if len(dto.MenuUrl) != 0 {
		byUrl, err1 := s.Dao.QueryMenuByMenuUrl(dto.MenuUrl)
		if err1 != nil {
			return err1
		}
		if byUrl != nil && item.Id != dto.Id {
			return errors.New("菜单URL已存在")
		}
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateMenu(dto)
}

// UpdateMenuStatus 更新菜单信息状态
func (s *MenuServiceImpl) UpdateMenuStatus(dto d.UpdateMenuStatusDto) error {
	return s.Dao.UpdateMenuStatus(dto)
}

// QueryMenuDetail 查询菜单信息详情
func (s *MenuServiceImpl) QueryMenuDetail(dto d.QueryMenuDetailDto) (*d.QueryMenuListDtoResp, error) {
	item, err := s.Dao.QueryMenuDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("菜单信息不存在")
	}

	return &d.QueryMenuListDtoResp{
		Id:         item.Id,                             // 主键
		MenuName:   item.MenuName,                       // 菜单名称
		MenuType:   item.MenuType,                       // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    item.Visible,                        // 显示状态（0:隐藏, 显示:1）
		Status:     item.Status,                         // 菜单状态(1:正常，0:禁用)
		Sort:       item.Sort,                           // 排序
		ParentId:   item.ParentId,                       // 父ID
		MenuUrl:    item.MenuUrl,                        // 路由路径
		ApiUrl:     item.ApiUrl,                         // 接口URL
		MenuIcon:   item.MenuIcon,                       // 菜单图标
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryMenuById 根据id查询菜单信息详情
func (s *MenuServiceImpl) QueryMenuById(id int64) (*d.QueryMenuListDtoResp, error) {
	item, err := s.Dao.QueryMenuById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("菜单信息不存在")
	}

	return &d.QueryMenuListDtoResp{
		Id:         item.Id,                             // 主键
		MenuName:   item.MenuName,                       // 菜单名称
		MenuType:   item.MenuType,                       // 菜单类型(1：目录   2：菜单   3：按钮)
		Visible:    item.Visible,                        // 显示状态（0:隐藏, 显示:1）
		Status:     item.Status,                         // 菜单状态(1:正常，0:禁用)
		Sort:       item.Sort,                           // 排序
		ParentId:   item.ParentId,                       // 父ID
		MenuUrl:    item.MenuUrl,                        // 路由路径
		ApiUrl:     item.ApiUrl,                         // 接口URL
		MenuIcon:   item.MenuIcon,                       // 菜单图标
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryMenuList 查询菜单信息列表
func (s *MenuServiceImpl) QueryMenuList(dto d.QueryMenuListDto) ([]*d.QueryMenuListDtoResp, error) {
	result, err := s.Dao.QueryMenuList(dto)

	if err != nil {
		return nil, err
	}

	var list []*d.QueryMenuListDtoResp

	for _, item := range result {
		resp := &d.QueryMenuListDtoResp{
			Id:         item.Id,                             // 主键
			MenuName:   item.MenuName,                       // 菜单名称
			MenuType:   item.MenuType,                       // 菜单类型(1：目录   2：菜单   3：按钮)
			Visible:    item.Visible,                        // 显示状态（0:隐藏, 显示:1）
			Status:     item.Status,                         // 菜单状态(1:正常，0:禁用)
			Sort:       item.Sort,                           // 排序
			ParentId:   item.ParentId,                       // 父ID
			MenuUrl:    item.MenuUrl,                        // 路由路径
			ApiUrl:     item.ApiUrl,                         // 接口URL
			MenuIcon:   item.MenuIcon,                       // 菜单图标
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, nil
}
