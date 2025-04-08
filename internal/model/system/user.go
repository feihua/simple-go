package system

import "time"

// User 用户信息
type User struct {
	Id            int64     `gorm:"column:id" json:"id"`                         // 主键
	Mobile        string    `gorm:"column:mobile" json:"mobile"`                 // 手机号码
	UserName      string    `gorm:"column:user_name" json:"userName"`            // 用户账号
	NickName      string    `gorm:"column:nick_name" json:"nickName"`            // 用户昵称
	UserType      string    `gorm:"column:user_type" json:"userType"`            // 用户类型（00系统用户）
	Avatar        string    `gorm:"column:avatar" json:"avatar"`                 // 头像路径
	Email         string    `gorm:"column:email" json:"email"`                   // 用户邮箱
	Password      string    `gorm:"column:password" json:"password"`             // 密码
	Status        int32     `gorm:"column:status" json:"status"`                 // 状态(1:正常，0:禁用)
	DeptId        int64     `gorm:"column:dept_id" json:"deptId"`                // 部门ID
	LoginIp       string    `gorm:"column:login_ip" json:"loginIp"`              // 最后登录IP
	LoginDate     time.Time `gorm:"column:login_date" json:"loginDate"`          // 最后登录时间
	LoginBrowser  string    `gorm:"column:login_browser" json:"loginBrowser"`    // 浏览器类型
	LoginOs       string    `gorm:"column:login_os" json:"loginOs"`              // 操作系统
	PwdUpdateDate time.Time `gorm:"column:pwd_update_date" json:"pwdUpdateDate"` // 密码最后更新时间
	Remark        string    `gorm:"column:remark" json:"remark"`                 // 备注
	DelFlag       int32     `gorm:"column:del_flag" json:"delFlag"`              // 删除标志（0代表删除 1代表存在）
	CreateBy      string    `gorm:"column:create_by" json:"createBy"`            // 创建者
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateBy      string    `gorm:"column:update_by" json:"updateBy"`            // 更新者
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

func (model User) TableName() string {
	return "sys_user"
}
