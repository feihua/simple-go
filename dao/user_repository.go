package dao

import (
	"fmt"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(DB *gorm.DB) *UserDao {
	return &UserDao{db: DB}
}

// CreateUser 创建用户
func (u UserDao) CreateUser(dto dto.UserDto) error {

	user := models.User{
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

func (u UserDao) QueryUserByUsername(username string) []models.User {
	var sysUser []models.User

	u.db.Find(&sysUser, u.db.Where("username = ?", username))

	return sysUser
}

// QueryUserList 查询用户列表
func (u UserDao) QueryUserList(current int, pageSize int) ([]models.User, int) {

	var total = 0
	var sysUser []models.User
	u.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&sysUser)

	u.db.Model(&models.User{}).Count(&total)

	for k, v := range sysUser {
		fmt.Println(k, v)
	}

	return sysUser, total
}

// UpdateUser 更新用户
func (u UserDao) UpdateUser(dto dto.UserDto) error {

	user := models.User{
		Id:       dto.Id,
		Mobile:   dto.Mobile,
		UserName: dto.UserName,
		Password: dto.Password,
		StatusId: dto.StatusId,
		Sort:     dto.Sort,
		Remark:   dto.Remark,
	}

	err := u.db.Model(&user).Update(&user).Error

	return err
}

// DeleteUserByIds 删除用户
func (u UserDao) DeleteUserByIds(ids []int64) error {
	return u.db.Where("id in (?)", ids).Delete(&models.User{}).Error
}
