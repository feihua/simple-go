package system

import (
	"time"
)

// SysUser 用户信息
type SysUser struct {
	Id            int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`                           // 主键
	Mobile        string    `gorm:"column:mobile;NOT NULL" json:"mobile"`                                     // 手机号码
	UserName      string    `gorm:"column:user_name;NOT NULL" json:"user_name"`                               // 用户账号
	NickName      string    `gorm:"column:nick_name;NOT NULL" json:"nick_name"`                               // 用户昵称
	UserType      string    `gorm:"column:user_type;default:00;NOT NULL" json:"user_type"`                    // 用户类型（00系统用户）
	Avatar        string    `gorm:"column:avatar;NOT NULL" json:"avatar"`                                     // 头像路径
	Email         string    `gorm:"column:email;NOT NULL" json:"email"`                                       // 用户邮箱
	Password      string    `gorm:"column:password;NOT NULL" json:"password"`                                 // 密码
	Status        int       `gorm:"column:status;default:1;NOT NULL" json:"status"`                           // 状态(1:正常，0:禁用)
	DeptId        int64     `gorm:"column:dept_id;default:1;NOT NULL" json:"dept_id"`                         // 部门ID
	LoginIp       string    `gorm:"column:login_ip;NOT NULL" json:"login_ip"`                                 // 最后登录IP
	LoginDate     time.Time `gorm:"column:login_date" json:"login_date"`                                      // 最后登录时间
	LoginBrowser  string    `gorm:"column:login_browser;NOT NULL" json:"login_browser"`                       // 浏览器类型
	LoginOs       string    `gorm:"column:login_os;NOT NULL" json:"login_os"`                                 // 操作系统
	PwdUpdateDate time.Time `gorm:"column:pwd_update_date" json:"pwd_update_date"`                            // 密码最后更新时间
	Remark        string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 备注
	DelFlag       int       `gorm:"column:del_flag;default:1;NOT NULL" json:"del_flag"`                       // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `gorm:"column:create_by;NOT NULL" json:"create_by"`                               // 创建者
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
	UpdateBy      string    `gorm:"column:update_by;NOT NULL" json:"update_by"`                               // 更新者
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`                                    // 更新时间
}

func (m *SysUser) TableName() string {
	return "sys_user"
}
