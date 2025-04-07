package role

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
)

// RoleServiceImpl 角色操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 13:47
*/
type RoleServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewRoleServiceImpl(dao *dao.DaoImpl) RoleService {
	return &RoleServiceImpl{Dao: dao}
}

// CreateRole 创建角色
func (r *RoleServiceImpl) CreateRole(dto dto.RoleDto) error {
	return r.Dao.RoleDao.CreateRole(dto)
}

// QueryRoleList 查询角色列表
func (r *RoleServiceImpl) QueryRoleList(role dto.QueryRoleListDto) ([]model.Role, int64) {
	return r.Dao.RoleDao.QueryRoleList(role)
}

// UpdateRole 更新角色
func (r *RoleServiceImpl) UpdateRole(roleDto dto.RoleDto) error {
	return r.Dao.RoleDao.UpdateRole(roleDto)
}

// DeleteRoleByIds 删除角色
func (r *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return r.Dao.RoleDao.DeleteRoleByIds(ids)
}

// QueryRoleMenuList 根据角色Id查询角色菜单
func (r *RoleServiceImpl) QueryRoleMenuList(roleId int64) (*dto.QueryRoleMenuListDtoResp, error) {
	list, err := r.Dao.MenuDao.QueryMenuList()
	if err != nil {
		return nil, errors.New("查询所有菜单异常")
	}
	if len(list) == 0 {
		return nil, errors.New("查询所有菜单异常")
	}

	var menuIds []int64
	var menuList []dto.MenuDto

	for _, menu := range list {
		menuIds = append(menuIds, menu.Id)
		menuList = append(menuList, dto.MenuDto{
			Id:       menu.Id,
			MenuName: menu.MenuName,
			MenuType: menu.MenuType,
			StatusId: menu.StatusId,
			Sort:     menu.Sort,
			ParentId: menu.ParentId,
			MenuUrl:  menu.MenuUrl,
			ApiUrl:   menu.ApiUrl,
			MenuIcon: menu.MenuIcon,
			Remark:   menu.Remark,
		})
	}

	if roleId != 1 {
		ids, err1 := r.Dao.RoleMenuDao.QueryRoleMenuList(roleId)
		if err1 != nil {
			return nil, errors.New(err1.Error())
		}
		menuIds = ids
	}

	return &dto.QueryRoleMenuListDtoResp{
		MenuIds:  menuIds,
		MenuList: menuList,
	}, nil
}

// UpdateRoleMenuList 更新角色菜单
func (r *RoleServiceImpl) UpdateRoleMenuList(roleMenu dto.RoleMenuDtoRequest) error {
	if roleMenu.RoleId == 1 {
		return errors.New("不能修改超级管理员的权限")
	}
	role, err := r.Dao.RoleDao.QueryRoleById(roleMenu.RoleId)
	if err != nil {
		return errors.New("查询角色异常")
	}
	if role == nil {
		return errors.New("角色不存在")
	}
	if role.StatusId != 1 {
		return errors.New("角色已禁用")
	}
	if len(roleMenu.MenuIds) == 0 {
		return errors.New("菜单不能为空")
	}
	return r.Dao.RoleMenuDao.UpdateRoleMenuList(roleMenu.RoleId, roleMenu.MenuIds)
}
