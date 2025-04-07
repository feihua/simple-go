package dao

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
	"gorm.io/gorm"
)

// UserDao 用户操作
/*
Author: LiuFeiHua
Date: 2024/4/28 10:31
*/
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(DB *gorm.DB) *UserDao {
	return &UserDao{db: DB}
}

// CreateUser 创建用户
func (u UserDao) CreateUser(dto dto.UserDto) error {

	user := model.User{
		Mobile:   dto.Mobile,
		UserName: dto.UserName,
		Password: dto.Password,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := u.db.Create(&user).Error

	return err
}

// QueryUserByUserId 根据用户id查询用户
func (u UserDao) QueryUserByUserId(userId int64) (*model.User, error) {
	var sysUser model.User

	err := u.db.First(&sysUser, u.db.Where("id = ?", userId)).Error

	switch {
	case err == nil:
		return &sysUser, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserByUsername 根据用户名称查询用户
func (u UserDao) QueryUserByUsername(username string) (*model.User, error) {
	var sysUser model.User

	err := u.db.First(&sysUser, u.db.Where("user_name = ?", username)).Error

	switch {
	case err == nil:
		return &sysUser, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserByUsernameOrMobile 根据用户名称或者收集查询用户信息
func (u UserDao) QueryUserByUsernameOrMobile(account string) ([]model.User, error) {
	var sysUser []model.User

	err := u.db.Model(&model.User{}).Where("user_name = ? or mobile = ?", account, account).Find(&sysUser).Error

	return sysUser, err
}

// QueryUserApiUrl 根据用户id查询用户权限
func (u UserDao) QueryUserApiUrl(userId int64) ([]string, error) {
	sql := `select u.api_url
from sys_user_role t
         left join sys_role usr on t.role_id = usr.id
         left join sys_role_menu srm on usr.id = srm.role_id
         left join sys_menu u on srm.menu_id = u.id
where t.user_id = ?`

	var apiUrls []string
	u.db.Raw(sql, userId).Scan(&apiUrls)

	if apiUrls == nil {
		return nil, errors.New("用户还没有分配权限")
	}

	return apiUrls, nil
}

// QueryUserMenus 根据用户id查询用户权限
func (u UserDao) QueryUserMenus(userId int64) ([]model.Menu, error) {
	sql := `select u.*
from sys_user_role t
         left join sys_role usr on t.role_id = usr.id
         left join sys_role_menu srm on usr.id = srm.role_id
         left join sys_menu u on srm.menu_id = u.id
where t.user_id = ?`

	var list []model.Menu
	u.db.Raw(sql, userId).Scan(&list)

	return list, nil
}

// QueryUserList 查询用户列表
func (u UserDao) QueryUserList(userListDto dto.QueryUserListDto) ([]model.User, int64) {

	pageNo := userListDto.PageNo
	pageSize := userListDto.PageSize
	var total int64 = 0
	var sysUser []model.User

	tx := u.db.Model(&model.User{})

	if userListDto.Mobile != "" {
		tx.Where("mobile=?", userListDto.Mobile)
	}
	if userListDto.UserName != "" {
		tx.Where("user_name=?", userListDto.UserName)
	}
	if userListDto.StatusId != 2 {
		tx.Where("status_id=?", userListDto.StatusId)
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&sysUser)
	tx.Count(&total)

	return sysUser, total
}

// UpdateUser 更新用户
func (u UserDao) UpdateUser(dto dto.UserDto) error {

	user := model.User{
		Id:       dto.Id,
		Mobile:   dto.Mobile,
		UserName: dto.UserName,
		Password: dto.Password,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := u.db.Model(&user).Updates(&user).Error

	return err
}

// DeleteUserByIds 删除用户
func (u UserDao) DeleteUserByIds(ids []int64) error {
	return u.db.Where("id in (?)", ids).Delete(&model.User{}).Error
}
