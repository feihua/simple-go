package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/redis"
	"github.com/feihua/simple-go/pkg/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// UserServiceImpl 用户信息操作实现
type UserServiceImpl struct {
	userDao     *system.UserDao
	userRoleDao *system.UserRoleDao
	MenuDao     *system.MenuDao
	RoleDao     *system.RoleDao
	UserPostDao *system.UserPostDao
}

func NewUserServiceImpl(dao *system.UserDao, userRoleDao *system.UserRoleDao, MenuDao *system.MenuDao, RoleDao *system.RoleDao, UserPostDao *system.UserPostDao) UserService {
	return &UserServiceImpl{
		userDao:     dao,
		userRoleDao: userRoleDao,
		MenuDao:     MenuDao,
		RoleDao:     RoleDao,
	}
}

// CreateUser 添加用户信息
func (s *UserServiceImpl) CreateUser(dto d.AddUserDto) error {
	byName, err := s.userDao.QueryUserByName(dto.UserName)
	if err != nil {
		return err
	}
	if byName != nil {
		return errors.New("用户名已存在")
	}
	byMobile, err := s.userDao.QueryUserByMobile(dto.Mobile)
	if err != nil {
		return err
	}
	if byMobile != nil {
		return errors.New("手机号码已存在")
	}
	byEmail, err := s.userDao.QueryUserByEmail(dto.Email)
	if err != nil {
		return err
	}
	if byEmail != nil {
		return errors.New("邮箱已存在")
	}
	avatar := dto.Avatar
	if len(avatar) == 0 {
		avatar = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	}
	dto.Avatar = avatar

	id, err := s.userDao.CreateUser(dto)
	if err != nil {
		return err
	}

	err = s.UserPostDao.DeleteUserPostByUserId(id)
	if err != nil {
		return err
	}

	for _, postId := range dto.PostIds {
		err = s.UserPostDao.CreateUserPost(d.AddUserPostDto{
			UserId: id,
			PostId: postId,
		})
		if err != nil {
			return err
		}
	}

	return s.userRoleDao.DeleteUserRoleByUserId(id)
}

// DeleteUserByIds 删除用户信息
func (s *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	for _, id := range ids {
		b := s.userRoleDao.IsAdministrator(id)
		if b {
			return errors.New("不允许操作超级管理员用户")
		}
	}

	err := s.UserPostDao.DeleteUserPostByUserIds(ids)
	if err != nil {
		return err
	}

	err = s.userRoleDao.DeleteUserRoleByUserIds(ids)
	if err != nil {
		return err
	}

	return s.userDao.DeleteUserByIds(ids)
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(dto d.UpdateUserDto) error {
	b := s.userRoleDao.IsAdministrator(dto.Id)
	if b {
		return errors.New("不允许操作超级管理员用户")
	}

	item, err := s.userDao.QueryUserById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("用户信息不存在")
	}

	byName, err := s.userDao.QueryUserByName(dto.UserName)
	if err != nil {
		return err
	}
	if byName != nil && item.Id != dto.Id {
		return errors.New("用户名已存在")
	}
	byMobile, err := s.userDao.QueryUserByMobile(dto.Mobile)
	if err != nil {
		return err
	}
	if byMobile != nil && item.Id != dto.Id {
		return errors.New("手机号码已存在")
	}
	byEmail, err := s.userDao.QueryUserByEmail(dto.Email)
	if err != nil {
		return err
	}
	if byEmail != nil && item.Id != dto.Id {
		return errors.New("邮箱已存在")
	}
	avatar := dto.Avatar
	if len(avatar) == 0 {
		avatar = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	}
	dto.Avatar = avatar

	err = s.UserPostDao.DeleteUserPostByUserId(dto.Id)
	if err != nil {
		return err
	}

	for _, postId := range dto.PostIds {
		err = s.UserPostDao.CreateUserPost(d.AddUserPostDto{
			UserId: dto.Id,
			PostId: postId,
		})
		if err != nil {
			return err
		}
	}

	dto.LoginIp = item.LoginIp             // 最后登录IP
	dto.LoginDate = item.LoginDate         // 最后登录时间
	dto.LoginBrowser = item.LoginBrowser   // 浏览器类型
	dto.LoginOs = item.LoginOs             // 操作系统
	dto.PwdUpdateDate = item.PwdUpdateDate // 密码最后更新时间
	dto.Remark = item.Remark               // 备注
	dto.DelFlag = item.DelFlag             // 删除标志（0代表删除 1代表存在）
	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.userDao.UpdateUser(dto)
}

// UpdateUserStatus 更新用户信息状态
func (s *UserServiceImpl) UpdateUserStatus(dto d.UpdateUserStatusDto) error {
	for _, id := range dto.Ids {
		b := s.userRoleDao.IsAdministrator(id)
		if b {
			return errors.New("不允许操作超级管理员用户")
		}
	}
	return s.userDao.UpdateUserStatus(dto)
}

// QueryUserDetail 查询用户信息详情
func (s *UserServiceImpl) QueryUserDetail(dto d.QueryUserDetailDto) (*d.QueryUserListDtoResp, error) {
	item, err := s.userDao.QueryUserDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("用户信息不存在")
	}

	return &d.QueryUserListDtoResp{
		Id:            item.Id,                                // 主键
		Mobile:        item.Mobile,                            // 手机号码
		UserName:      item.UserName,                          // 用户账号
		NickName:      item.NickName,                          // 用户昵称
		UserType:      item.UserType,                          // 用户类型（00系统用户）
		Avatar:        item.Avatar,                            // 头像路径
		Email:         item.Email,                             // 用户邮箱
		Password:      item.Password,                          // 密码
		Status:        item.Status,                            // 状态(1:正常，0:禁用)
		DeptId:        item.DeptId,                            // 部门ID
		LoginIp:       item.LoginIp,                           // 最后登录IP
		LoginDate:     utils.TimeToString(item.LoginDate),     // 最后登录时间
		LoginBrowser:  item.LoginBrowser,                      // 浏览器类型
		LoginOs:       item.LoginOs,                           // 操作系统
		PwdUpdateDate: utils.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
		Remark:        item.Remark,                            // 备注
		DelFlag:       item.DelFlag,                           // 删除标志（0代表删除 1代表存在）
		CreateBy:      item.CreateBy,                          // 创建者
		CreateTime:    utils.TimeToStr(item.CreateTime),       // 创建时间
		UpdateBy:      item.UpdateBy,                          // 更新者
		UpdateTime:    utils.TimeToString(item.UpdateTime),    // 更新时间
	}, nil
}

// QueryUserList 查询用户信息列表
func (s *UserServiceImpl) QueryUserList(dto d.QueryUserListDto) ([]*d.QueryUserListDtoResp, int64, error) {
	result, i, err := s.userDao.QueryUserList(dto)

	if err != nil {
		return nil, 0, err
	}

	var list []*d.QueryUserListDtoResp

	for _, item := range result {
		resp := &d.QueryUserListDtoResp{
			Id:            item.Id,                                // 主键
			Mobile:        item.Mobile,                            // 手机号码
			UserName:      item.UserName,                          // 用户账号
			NickName:      item.NickName,                          // 用户昵称
			UserType:      item.UserType,                          // 用户类型（00系统用户）
			Avatar:        item.Avatar,                            // 头像路径
			Email:         item.Email,                             // 用户邮箱
			Password:      item.Password,                          // 密码
			Status:        item.Status,                            // 状态(1:正常，0:禁用)
			DeptId:        item.DeptId,                            // 部门ID
			LoginIp:       item.LoginIp,                           // 最后登录IP
			LoginDate:     utils.TimeToString(item.LoginDate),     // 最后登录时间
			LoginBrowser:  item.LoginBrowser,                      // 浏览器类型
			LoginOs:       item.LoginOs,                           // 操作系统
			PwdUpdateDate: utils.TimeToString(item.PwdUpdateDate), // 密码最后更新时间
			Remark:        item.Remark,                            // 备注
			DelFlag:       item.DelFlag,                           // 删除标志（0代表删除 1代表存在）
			CreateBy:      item.CreateBy,                          // 创建者
			CreateTime:    utils.TimeToStr(item.CreateTime),       // 创建时间
			UpdateBy:      item.UpdateBy,                          // 更新者
			UpdateTime:    utils.TimeToString(item.UpdateTime),    // 更新时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}

// Login 登录
func (u *UserServiceImpl) Login(loginDto d.LoginDto) (*string, error) {
	utils.Logger.Debugf("登录参数: %+v", loginDto)
	user, err := u.userDao.QueryUserByAccount(loginDto.Account)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		utils.Logger.Debugf("用户不存在, 请求参数：%+v, 异常信息: %s", loginDto, err.Error())
		return nil, errors.New("用户不存在")
	case err != nil:
		utils.Logger.Debugf("查询用户异常, 请求参数：%+v, 异常信息: %s", loginDto, err.Error())
		return nil, errors.New("查询用户表异常")
	}

	if user.Password != loginDto.Password {
		return nil, errors.New("密码错误")
	}

	if user.Status != 1 {
		return nil, errors.New("用户被禁用")
	}

	var apiUrl []string
	if u.userRoleDao.IsAdministrator(user.Id) {
		apiUrl, err = u.MenuDao.QueryApiUrlList()
	} else {
		apiUrl, err = u.userDao.QueryApiUrls(user.Id)
	}

	if err != nil {
		return nil, errors.New("获取用户权限失败")
	}

	if apiUrl == nil {
		return nil, errors.New("用户还没有分配权限")
	}

	now := time.Now().Unix()
	jwtInfo := config.GlobalAppConfig.Jwt
	token, err := createJwtToken(jwtInfo.AccessSecret, now, jwtInfo.AccessExpire, user.Id, user.UserName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	apiUrlStr := strings.Join(apiUrl, ",")
	ctx := context.Background()
	redis.Rdb.Set(ctx, "simple:apiUrl:"+strconv.FormatInt(user.Id, 10), apiUrlStr, time.Duration(jwtInfo.AccessExpire)*time.Second)
	return &token, nil

}

// 生成jwt的token
func createJwtToken(secretKey string, iat, seconds, userId int64, userName string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["userName"] = userName
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// QueryUserMenu 查询用户菜单权限信息
func (u *UserServiceImpl) QueryUserMenu(userId int64, userName string) (*d.QueryUserMenuDtoResp, error) {
	user, err := u.userDao.QueryUserById(userId)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		utils.Logger.Debugf("用户不存在, 请求参数userId：%+v, 异常信息: %s", userId, err.Error())
		return nil, errors.New("用户不存在")
	case err != nil:
		utils.Logger.Debugf("查询用户异常, 请求参数userId：%+v, 异常信息: %s", userId, err.Error())
		return nil, errors.New("查询用户表异常")
	}

	if user.Status != 1 {
		return nil, errors.New("用户被禁用")
	}

	var menuList []*m.Menu
	if u.userRoleDao.IsAdministrator(userId) {
		menuList, err = u.MenuDao.QueryAllMenuList()
	} else {
		menuList, err = u.userDao.QueryUserMenus(userId)
	}

	if err != nil {
		return nil, errors.New("查询菜单异常")
	}

	var list []d.UserMenuDto
	var apiUrl []string
	for _, menu := range menuList {
		if menu.ApiUrl != "" {
			apiUrl = append(apiUrl, menu.ApiUrl)
		}
		if menu.MenuType != 3 {
			list = append(list, d.UserMenuDto{
				Id:       menu.Id,
				MenuName: menu.MenuName,
				ParentId: menu.ParentId,
				MenuUrl:  menu.MenuUrl,
				MenuIcon: menu.MenuIcon,
			})
		}
	}

	avatar := "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	return &d.QueryUserMenuDtoResp{
		Id:       userId,
		UserName: userName,
		Avatar:   avatar,
		Menus:    list,
		ApiUrls:  apiUrl,
	}, nil
}

func (u *UserServiceImpl) QueryUserRoleList(dto d.QueryUserRoleListDto) (*d.QueryUserRoleListDataDtoResp, error) {
	byId, err := u.userDao.QueryUserById(dto.UserId)
	if err != nil {
		return nil, err
	}
	if byId == nil {
		return nil, errors.New("用户不存在")
	}

	roleList, _, err := u.RoleDao.QueryRoleList(d.QueryRoleListDto{
		PageNo:   dto.PageNo,
		PageSize: dto.PageSize,
		// RoleName:  "",
		// RoleKey:   "",
		// DataScope: 0,
		// Status:    0,
	})
	if err != nil {
		return nil, err
	}

	var roleIds []int64
	var list []d.QueryRoleListDtoResp
	for _, item := range roleList {
		roleIds = append(roleIds, item.Id)
		list = append(list, d.QueryRoleListDtoResp{
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
		})
	}

	b := u.userRoleDao.IsAdministrator(dto.UserId)
	if !b {
		roleIds, err = u.userRoleDao.QueryUserRoleIds(dto.UserId)
		if err != nil {
			return nil, err
		}
	}
	return &d.QueryUserRoleListDataDtoResp{
		RoleList: list,
		RoleIds:  roleIds,
	}, nil
}

func (s *UserServiceImpl) UpdateUserRole(dto d.UpdateUserRoleDto) error {
	b := s.userRoleDao.IsAdministrator(dto.UserId)
	if b {
		return errors.New("不允许操作超级管理员用户")
	}

	err := s.userRoleDao.DeleteUserRoleByUserId(dto.UserId)
	if err != nil {
		return err
	}

	for _, id := range dto.RoleIds {
		err = s.userRoleDao.CreateUserRole(d.AddUserRoleDto{
			UserId: dto.UserId,
			RoleId: id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
