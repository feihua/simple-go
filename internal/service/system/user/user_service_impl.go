package user

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
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
}

func NewUserServiceImpl(dao *system.UserDao, userRoleDao *system.UserRoleDao, MenuDao *system.MenuDao, RoleDao *system.RoleDao) UserService {
	return &UserServiceImpl{
		userDao:     dao,
		userRoleDao: userRoleDao,
		MenuDao:     MenuDao,
		RoleDao:     RoleDao,
	}
}

// CreateUser 添加用户信息
func (s *UserServiceImpl) CreateUser(dto a.AddUserDto) error {
	return s.userDao.CreateUser(dto)
}

// DeleteUserByIds 删除用户信息
func (s *UserServiceImpl) DeleteUserByIds(ids []int64) error {
	return s.userDao.DeleteUserByIds(ids)
}

// UpdateUser 更新用户信息
func (s *UserServiceImpl) UpdateUser(dto a.UpdateUserDto) error {
	return s.userDao.UpdateUser(dto)
}

// UpdateUserStatus 更新用户信息状态
func (s *UserServiceImpl) UpdateUserStatus(dto a.UpdateUserStatusDto) error {
	return s.userDao.UpdateUserStatus(dto)
}

// QueryUserDetail 查询用户信息详情
func (s *UserServiceImpl) QueryUserDetail(dto a.QueryUserDetailDto) (b.User, error) {
	return s.userDao.QueryUserDetail(dto)
}

// QueryUserList 查询用户信息列表
func (s *UserServiceImpl) QueryUserList(dto a.QueryUserListDto) ([]b.User, int64) {
	return s.userDao.QueryUserList(dto)
}

// Login 登录
func (u *UserServiceImpl) Login(loginDto a.LoginDto) (*a.LoginDtoResp, error) {
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
	return &a.LoginDtoResp{
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
