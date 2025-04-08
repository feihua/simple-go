package menu

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
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
func (s *MenuServiceImpl) CreateMenu(dto a.AddMenuDto) error {
	return s.Dao.CreateMenu(dto)
}

// DeleteMenuByIds 删除菜单信息
func (s *MenuServiceImpl) DeleteMenuByIds(ids []int64) error {
	return s.Dao.DeleteMenuByIds(ids)
}

// UpdateMenu 更新菜单信息
func (s *MenuServiceImpl) UpdateMenu(dto a.UpdateMenuDto) error {
	return s.Dao.UpdateMenu(dto)
}

// UpdateMenuStatus 更新菜单信息状态
func (s *MenuServiceImpl) UpdateMenuStatus(dto a.UpdateMenuStatusDto) error {
	return s.Dao.UpdateMenuStatus(dto)
}

// QueryMenuDetail 查询菜单信息详情
func (s *MenuServiceImpl) QueryMenuDetail(dto a.QueryMenuDetailDto) (b.Menu, error) {
	return s.Dao.QueryMenuDetail(dto)
}

// QueryMenuList 查询菜单信息列表
func (s *MenuServiceImpl) QueryMenuList(dto a.QueryMenuListDto) ([]b.Menu, int64) {
	return s.Dao.QueryMenuList(dto)
}
