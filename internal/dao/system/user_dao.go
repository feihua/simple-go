package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(DB *gorm.DB) *UserDao {
	return &UserDao{
		db: DB,
	}
}

// CreateUser 添加用户信息
func (b UserDao) CreateUser(dto system.AddUserDto) (int64, error) {
	item := m.User{
		Mobile:   dto.Mobile,   // 手机号码
		UserName: dto.UserName, // 用户账号
		NickName: dto.NickName, // 用户昵称
		UserType: dto.UserType, // 用户类型（00系统用户）
		Avatar:   dto.Avatar,   // 头像路径
		Email:    dto.Email,    // 用户邮箱
		Password: dto.Password, // 密码
		Status:   dto.Status,   // 状态(1:正常，0:禁用)
		DeptId:   dto.DeptId,   // 部门ID
		Remark:   dto.Remark,   // 备注
		DelFlag:  1,            // 删除标志（0代表删除 1代表存在）
		CreateBy: dto.CreateBy, // 创建者

	}

	err := b.db.Create(&item).Error
	if err != nil {
		return 0, err
	}
	return item.Id, err
}

// DeleteUserByIds 根据id删除用户信息
func (b UserDao) DeleteUserByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.User{}).Error
}

// UpdateUser 更新用户信息
func (b UserDao) UpdateUser(dto system.UpdateUserDto) error {

	item := m.User{
		Id:            dto.Id,            // 主键
		Mobile:        dto.Mobile,        // 手机号码
		UserName:      dto.UserName,      // 用户账号
		NickName:      dto.NickName,      // 用户昵称
		UserType:      dto.UserType,      // 用户类型（00系统用户）
		Avatar:        dto.Avatar,        // 头像路径
		Email:         dto.Email,         // 用户邮箱
		Password:      dto.Password,      // 密码
		Status:        dto.Status,        // 状态(1:正常，0:禁用)
		DeptId:        dto.DeptId,        // 部门ID
		LoginIp:       dto.LoginIp,       // 最后登录IP
		LoginDate:     dto.LoginDate,     // 最后登录时间
		LoginBrowser:  dto.LoginBrowser,  // 浏览器类型
		LoginOs:       dto.LoginOs,       // 操作系统
		PwdUpdateDate: dto.PwdUpdateDate, // 密码最后更新时间
		Remark:        dto.Remark,        // 备注
		DelFlag:       dto.DelFlag,       // 删除标志（0代表删除 1代表存在）
		CreateBy:      dto.CreateBy,      // 创建者
		CreateTime:    dto.CreateTime,    // 创建时间
		UpdateBy:      dto.UpdateBy,      // 更新者
		UpdateTime:    &dto.UpdateTime,   // 更新时间
	}

	return b.db.Updates(&item).Error
}

// UpdateUserStatus 更新用户信息状态
func (b UserDao) UpdateUserStatus(dto system.UpdateUserStatusDto) error {

	return b.db.Model(&m.Dept{}).Where("id in (?)", dto.Ids).Update("status", dto.Status).Error
}

// QueryUserDetail 查询用户信息详情
func (b UserDao) QueryUserDetail(dto system.QueryUserDetailDto) (*m.User, error) {
	var item m.User
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserList 查询用户信息列表
func (b UserDao) QueryUserList(dto system.QueryUserListDto) ([]*m.User, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.User
	tx := b.db.Model(&m.User{})
	if len(dto.Mobile) > 0 {
		tx.Where("mobile like %?%", dto.Mobile) // 手机号码
	}
	if len(dto.UserName) > 0 {
		tx.Where("user_name like %?%", dto.UserName) // 用户账号
	}
	if len(dto.NickName) > 0 {
		tx.Where("nick_name like %?%", dto.NickName) // 用户昵称
	}

	if dto.Status != 2 {
		tx.Where("status=?", dto.Status) // 状态(1:正常，0:禁用)
	}
	if dto.DeptId != 0 {
		tx.Where("dept_id=?", dto.DeptId) // 部门ID
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}

// QueryUserByAccount 查询用户信息
func (b UserDao) QueryUserByAccount(account string) (*m.User, error) {
	var item m.User
	err := b.db.Where("mobile=? or user_name =?", account, account).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryApiUrls 根据用户id查询用户权限
func (u UserDao) QueryApiUrls(userId int64) ([]string, error) {
	sql := `select u.api_url
from sys_user_role t
         left join sys_role usr on t.role_id = usr.id
         left join sys_role_menu srm on usr.id = srm.role_id
         left join sys_menu u on srm.menu_id = u.id
where t.user_id = ?`

	var apiUrls []string
	err := u.db.Raw(sql, userId).Scan(&apiUrls).Error

	return apiUrls, err
}

// QueryUserById 根据id查询用户信息
func (b UserDao) QueryUserById(id int64) (*m.User, error) {
	var item m.User
	err := b.db.Where("id = ?", id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserByMobile 根据mobile查询用户
func (b UserDao) QueryUserByMobile(mobile string) (*m.User, error) {
	var item m.User
	err := b.db.Where("mobile = ?", mobile).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserByName 根据name查询用户
func (b UserDao) QueryUserByName(userName string) (*m.User, error) {
	var item m.User
	err := b.db.Where("user_name = ?", userName).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserByEmail 根据email查询用户
func (b UserDao) QueryUserByEmail(email string) (*m.User, error) {
	var item m.User
	err := b.db.Where("email = ?", email).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryUserMenus 根据用户id查询用户权限
func (u UserDao) QueryUserMenus(userId int64) ([]*m.Menu, error) {
	sql := `select u.*
			from sys_user_role t
					 left join sys_role usr on t.role_id = usr.id
					 left join sys_role_menu srm on usr.id = srm.role_id
					 left join sys_menu u on srm.menu_id = u.id
			where t.user_id = ?
`

	var list []*m.Menu
	err := u.db.Raw(sql, userId).Scan(&list).Error

	return list, err
}

// QueryUserByUserIds 根据userIds查询用户
func (b UserDao) QueryUserByUserIds(dto system.QueryRoleUserListDto, userIs []int64) ([]*m.User, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.User
	tx := b.db.Model(&m.User{}).Where("id in (?)", userIs)

	if len(dto.Mobile) > 0 {
		tx.Where("mobile like %?%", dto.Mobile) // 手机号码
	}
	if len(dto.UserName) > 0 {
		tx.Where("user_name like %?%", dto.UserName) // 用户账号
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}

// QueryUserNotUserIds 排除userIds的用户
func (b UserDao) QueryUserNotUserIds(dto system.QueryRoleUserListDto, userIs []int64) ([]*m.User, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.User
	tx := b.db.Model(&m.User{}).Where("id not in (?)", userIs)

	if len(dto.Mobile) > 0 {
		tx.Where("mobile like %?%", dto.Mobile) // 手机号码
	}
	if len(dto.UserName) > 0 {
		tx.Where("user_name like %?%", dto.UserName) // 用户账号
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}

// ExistUserByDeptId 查询部门是否存在用户
func (b UserDao) ExistUserByDeptId(deptId int64) ([]*m.User, error) {
	var item []*m.User
	err := b.db.Model(&m.Dept{}).Where("dept_id = ?", deptId).Scan(&item).Error

	return item, err
}
