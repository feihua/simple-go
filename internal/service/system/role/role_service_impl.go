package role

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"strconv"
	"time"
)

// RoleServiceImpl 角色信息操作实现
type RoleServiceImpl struct {
	Dao         *system.RoleDao
	userRoleDao *system.UserRoleDao
	RoleMenuDao *system.RoleMenuDao
	MenuDao     *system.MenuDao
	UserDao     *system.UserDao
}

func NewRoleServiceImpl(dao *system.RoleDao, userRoleDao *system.UserRoleDao, RoleMenuDao *system.RoleMenuDao, MenuDao *system.MenuDao, UserDao *system.UserDao) RoleService {
	return &RoleServiceImpl{
		Dao:         dao,
		userRoleDao: userRoleDao,
		RoleMenuDao: RoleMenuDao,
		MenuDao:     MenuDao,
		UserDao:     UserDao,
	}
}

// CreateRole 添加角色信息
func (s *RoleServiceImpl) CreateRole(dto d.AddRoleDto) error {
	byName, err := s.Dao.QueryRoleByName(dto.RoleName)
	if err != nil {
		return err
	}
	if byName != nil {
		return errors.New("角色名称已存在")
	}
	byKey, err := s.Dao.QueryRoleByKey(dto.RoleKey)
	if err != nil {
		return err
	}
	if byKey != nil {
		return errors.New("角色权限字符串已存在")
	}
	return s.Dao.CreateRole(dto)
}

// DeleteRoleByIds 删除角色信息
func (s *RoleServiceImpl) DeleteRoleByIds(ids []int64) error {
	return s.Dao.DeleteRoleByIds(ids)
}

// UpdateRole 更新角色信息
func (s *RoleServiceImpl) UpdateRole(dto d.UpdateRoleDto) error {
	if dto.Id == 1 {
		return errors.New("不允许操作超级管理员角色")
	}
	item, err := s.Dao.QueryRoleById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("角色信息不存在")
	}

	byName, err := s.Dao.QueryRoleByName(dto.RoleName)
	if err != nil {
		return err
	}
	if byName != nil && item.Id != dto.Id {
		return errors.New("角色名称已存在")
	}
	byKey, err := s.Dao.QueryRoleByKey(dto.RoleKey)
	if err != nil {
		return err
	}
	if byKey != nil && item.Id != dto.Id {
		return errors.New("角色权限字符串已存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdateRole(dto)
}

// UpdateRoleStatus 更新角色信息状态
func (s *RoleServiceImpl) UpdateRoleStatus(dto d.UpdateRoleStatusDto) error {
	for _, id := range dto.Ids {
		if id == 1 {
			return errors.New("不允许操作超级管理员角色")
		}
		byId, err := s.Dao.QueryRoleById(id)
		if err != nil {
			return err
		}
		if byId == nil {
			return errors.New("角色不存在")
		}
	}
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

// QueryAllocatedList 查询已分配用户角色列表
func (s *RoleServiceImpl) QueryAllocatedList(dto d.QueryRoleUserListDto) ([]*d.QueryUserListDtoResp, int64, error) {
	if dto.RoleId == 1 {
		return nil, 0, errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return nil, 0, err
	}

	if roleById == nil {
		return nil, 0, errors.New("角色信息不存在")
	}

	userIds, err := s.userRoleDao.QueryUserIds(dto.RoleId)
	if err != nil {
		return nil, 0, err
	}

	users, count, err := s.UserDao.QueryUserByUserIds(dto, userIds)

	if err != nil {
		return nil, 0, err
	}

	var list = make([]*d.QueryUserListDtoResp, 0, len(users))

	for _, item := range users {
		list = append(list, &d.QueryUserListDtoResp{
			Id:         item.Id,                             // 主键
			Mobile:     item.Mobile,                         // 手机号码
			UserName:   item.UserName,                       // 用户账号
			NickName:   item.NickName,                       // 用户昵称
			Avatar:     item.Avatar,                         // 头像路径
			Email:      item.Email,                          // 用户邮箱
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			DeptId:     item.DeptId,                         // 部门ID
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}
	return list, count, nil
}

// QueryUnallocatedList 查询未分配用户角色列表
func (s *RoleServiceImpl) QueryUnallocatedList(dto d.QueryRoleUserListDto) ([]*d.QueryUserListDtoResp, int64, error) {
	if dto.RoleId == 1 {
		return nil, 0, errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return nil, 0, err
	}

	if roleById == nil {
		return nil, 0, errors.New("角色信息不存在")
	}

	userIds, err := s.userRoleDao.QueryUserIds(dto.RoleId)
	if err != nil {
		return nil, 0, err
	}

	users, count, err := s.UserDao.QueryUserNotUserIds(dto, userIds)

	if err != nil {
		return nil, 0, err
	}

	var list = make([]*d.QueryUserListDtoResp, 0, len(users))

	for _, item := range users {
		list = append(list, &d.QueryUserListDtoResp{
			Id:         item.Id,                             // 主键
			Mobile:     item.Mobile,                         // 手机号码
			UserName:   item.UserName,                       // 用户账号
			NickName:   item.NickName,                       // 用户昵称
			Avatar:     item.Avatar,                         // 头像路径
			Email:      item.Email,                          // 用户邮箱
			Status:     item.Status,                         // 状态(1:正常，0:禁用)
			DeptId:     item.DeptId,                         // 部门ID
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间

		})
	}
	return list, count, nil
}

// CancelAuthUser 取消授权用户
func (s *RoleServiceImpl) CancelAuthUser(dto d.AuthUserDto) error {
	if dto.RoleId == 1 {
		return errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return err
	}

	if roleById == nil {
		return errors.New("角色信息不存在")
	}

	return s.Dao.CancelAuthUser(dto.UserId, dto.RoleId)
}

// BatchCancelAuthUser 批量取消授权用户
func (s *RoleServiceImpl) BatchCancelAuthUser(dto d.BatchAuthUserDto) error {
	if dto.RoleId == 1 {
		return errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return err
	}

	if roleById == nil {
		return errors.New("角色信息不存在")
	}

	return s.Dao.BatchCancelAuthUser(dto.UserIds, dto.RoleId)
}

// BatchAuthUser 批量选择用户授权
func (s *RoleServiceImpl) BatchAuthUser(dto d.BatchAuthUserDto) error {
	if dto.RoleId == 1 {
		return errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return err
	}

	if roleById == nil {
		return errors.New("角色信息不存在")
	}

	return s.Dao.BatchAuthUser(dto)
}

// QueryRoleMenuList 分页查询角色菜单关联列表
func (s *RoleServiceImpl) QueryRoleMenuList(dto d.QueryRoleMenuListDto) (*d.QueryRoleMenuListDataDtoResp, error) {
	allMenuList, err := s.MenuDao.QueryAllMenuList()
	if err != nil {
		return nil, err
	}

	var menuList []d.RoleMenuListDataDto
	var menuIds []int64

	for _, menu := range allMenuList {
		menuList = append(menuList, d.RoleMenuListDataDto{
			Key:           strconv.FormatInt(menu.Id, 10),
			Title:         menu.MenuName,
			ParentId:      menu.ParentId,
			Id:            menu.Id,
			Label:         menu.MenuName,
			IsPenultimate: menu.ParentId == 2,
		})
		menuIds = append(menuIds, menu.Id)
	}

	// 2.如果角色不是admin则根据roleId查询菜单
	if dto.RoleId != 1 {
		menuIds, err = s.RoleMenuDao.QueryMenuIds(dto.RoleId)

	}

	data := &d.QueryRoleMenuListDataDtoResp{
		MenuList: menuList,
		MenuIds:  menuIds,
	}

	return data, nil

}

// UpdateRoleMenu 添加角色菜单关联
func (s *RoleServiceImpl) UpdateRoleMenu(dto d.UpdateRoleMenuDto) error {
	if dto.RoleId == 1 {
		return errors.New("不允许取消超级管理员角色")
	}

	roleById, err := s.Dao.QueryRoleById(dto.RoleId)

	if err != nil {
		return err
	}

	if roleById == nil {
		return errors.New("角色信息不存在")
	}

	err = s.RoleMenuDao.DeleteRoleMenuByRoleId(dto.RoleId)
	if err != nil {
		return err
	}

	return s.RoleMenuDao.CreateRoleMenuBatch(dto)
}
