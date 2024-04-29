package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/redis"
	"github.com/feihua/simple-go/pkg/utils"
	"strconv"
	"strings"
	"time"
)

// UserServiceImpl 用户操作接口实现
/*
Author: LiuFeiHua
Date: 2024/4/16 13:58
*/
type UserServiceImpl struct {
	Dao *dao.DaoImpl
}

func NewUserServiceImpl(Dao *dao.DaoImpl) UserService {
	return &UserServiceImpl{
		Dao: Dao,
	}
}

// Login 登录
func (u *UserServiceImpl) Login(loginDto dto.UserLoginDto) (*dto.LoginDtoResp, error) {
	utils.Logger.Debugf("登录参数: %+v", loginDto)
	users, err := u.Dao.UserDao.QueryUserByUsernameOrMobile(loginDto.Account)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len(users) == 0 {
		return nil, errors.New("用户不存在")
	}

	user := users[0]
	if user.Password != loginDto.Password {
		return nil, errors.New("密码错误")
	}

	if user.StatusId != 1 {
		return nil, errors.New("用户被禁用")
	}

	var apiUrl []string
	if u.Dao.UserRoleDao.IsAdministrator(user.Id) {
		menuList, _ := u.Dao.MenuDao.QueryMenuList()
		for _, menu := range menuList {
			if menu.ApiUrl != "" {
				apiUrl = append(apiUrl, menu.ApiUrl)
			}
		}
	} else {
		url, err1 := u.Dao.UserDao.QueryUserApiUrl(user.Id)
		if err1 != nil {
			return nil, errors.New(err1.Error())
		}
		apiUrl = url
	}

	now := time.Now().Unix()
	token, err := createJwtToken(config.TokenInfo.AccessSecret, now, config.TokenInfo.AccessExpire, user.Id, user.UserName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	apiUrlStr := strings.Join(apiUrl, ",")
	ctx := context.Background()
	redis.Rdb.Set(ctx, "simple:apiUrl:"+strconv.FormatInt(user.Id, 10), apiUrlStr, time.Duration(config.TokenInfo.AccessExpire)*time.Second)
	return &dto.LoginDtoResp{
		Id:       user.Id,
		UserName: user.UserName,
		Token:    token,
	}, nil

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
func (u *UserServiceImpl) QueryUserMenu(userId int64, userName string) (*dto.QueryUserMenuDtoResp, error) {
	user, err := u.Dao.UserDao.QueryUserByUserId(userId)
	if err != nil {
		return nil, errors.New("查询用户异常")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.StatusId != 1 {
		return nil, errors.New("用户被禁用")
	}

	var menuList []models.Menu
	if u.Dao.UserRoleDao.IsAdministrator(userId) {
		menuList, err = u.Dao.MenuDao.QueryMenuList()
	} else {
		menuList, err = u.Dao.UserDao.QueryUserMenus(userId)
	}

	if err != nil {
		return nil, errors.New("查询菜单异常")
	}

	var list []dto.UserMenuDto
	var apiUrl []string
	for _, menu := range menuList {
		if menu.ApiUrl != "" {
			apiUrl = append(apiUrl, menu.ApiUrl)
		}
		if menu.MenuType != 3 {
			list = append(list, dto.UserMenuDto{
				Id:       menu.Id,
				MenuName: menu.MenuName,
				Sort:     menu.Sort,
				ParentId: menu.ParentId,
				MenuUrl:  menu.MenuUrl,
				MenuIcon: menu.MenuIcon,
			})
		}
	}

	avatar := "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	return &dto.QueryUserMenuDtoResp{
		Id:       userId,
		UserName: userName,
		Avatar:   avatar,
		Menus:    list,
		ApiUrls:  apiUrl,
	}, nil
}

// CreateUser 创建用户
func (u *UserServiceImpl) CreateUser(dto dto.UserDto) error {
	user, err := u.Dao.UserDao.QueryUserByUsername(dto.UserName)
	if err != nil {
		return errors.New("根据用户名称查询用户异常")
	}
	if user != nil {
		return errors.New("用户已存在")
	}
	return u.Dao.UserDao.CreateUser(dto)
}

// QueryUserList 查询用户列表
func (u *UserServiceImpl) QueryUserList(userListDto dto.QueryUserListDto) ([]models.User, int64) {
	return u.Dao.UserDao.QueryUserList(userListDto)
}

// UpdateUser 更新用户
func (u *UserServiceImpl) UpdateUser(userDto dto.UserDto) error {
	if userDto.Id == 1 {
		return errors.New("超级管理员不能修改角色")
	}
	return u.Dao.UserDao.UpdateUser(userDto)
}

// DeleteUserByIds 删除用户
func (u *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	return u.Dao.UserDao.DeleteUserByIds(ids)
}

// QueryUserRoleList 查询用户与角色关糸
func (u *UserServiceImpl) QueryUserRoleList(userId int64) ([]int64, error) {
	return u.Dao.UserRoleDao.QueryUserRoleList(userId)
}

// UpdateUserRoleList 更新用户与角色关糸
func (u *UserServiceImpl) UpdateUserRoleList(req dto.UpdateUserRoleDtoRequest) error {
	userId := req.UserId
	if u.Dao.UserRoleDao.IsAdministrator(userId) {
		return errors.New("超级管理员不能修改角色")
	}
	user, err := u.Dao.UserDao.QueryUserByUserId(userId)
	if err != nil {
		return errors.New("查询用户异常")
	}
	if user == nil {
		return errors.New("用户不存在")
	}
	if user.StatusId != 1 {
		return errors.New("用户被禁用")
	}

	if len(req.RoleId) == 0 {
		return errors.New("角色不能为空")
	}

	var userRoles []models.UserRole
	for _, roleId := range req.RoleId {
		userRoles = append(userRoles, models.UserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}

	return u.Dao.UserRoleDao.UpdateUserRoleList(userId, userRoles)
}
