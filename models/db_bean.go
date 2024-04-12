package models

import (
	"time"
)

type SysConfig struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Value          string    `json:"value" xorm:"not null comment('数据值') VARCHAR(100) 'value'"`
	Label          string    `json:"label" xorm:"not null comment('标签名') VARCHAR(100) 'label'"`
	Type           string    `json:"type" xorm:"not null comment('类型') VARCHAR(100) 'type'"`
	Description    string    `json:"description" xorm:"not null comment('描述') VARCHAR(100) 'description'"`
	Sort           string    `json:"sort" xorm:"not null comment('排序（升序）') DECIMAL 'sort'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	Remarks        string    `json:"remarks" xorm:"comment('备注信息') VARCHAR(255) 'remarks'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysDept struct {
	Id             int       `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name           string    `json:"name" xorm:"comment('机构名称') VARCHAR(50) 'name'"`
	ParentId       int       `json:"parent_id" xorm:"comment('上级机构ID，一级机构为0') BIGINT(20) 'parent_id'"`
	OrderNum       int64     `json:"order_num" xorm:"comment('排序') INT(11) 'order_num'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysDict struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Value          string    `json:"value" xorm:"not null comment('数据值') VARCHAR(100) 'value'"`
	Label          string    `json:"label" xorm:"not null comment('标签名') VARCHAR(100) 'label'"`
	Type           string    `json:"type" xorm:"not null comment('类型') VARCHAR(100) 'type'"`
	Description    string    `json:"description" xorm:"not null comment('描述') VARCHAR(100) 'description'"`
	Sort           int64     `json:"sort" xorm:"not null comment('排序（升序）') DECIMAL 'sort'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	Remarks        string    `json:"remarks" xorm:"comment('备注信息') VARCHAR(255) 'remarks'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysLog struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	UserName       string    `json:"user_name" xorm:"comment('用户名') VARCHAR(50) 'user_name'"`
	Operation      string    `json:"operation" xorm:"comment('用户操作') VARCHAR(50) 'operation'"`
	Method         string    `json:"method" xorm:"comment('请求方法') VARCHAR(200) 'method'"`
	Params         string    `json:"params" xorm:"comment('请求参数') VARCHAR(5000) 'params'"`
	Time           int64     `json:"time" xorm:"not null comment('执行时长(毫秒)') BIGINT(20) 'time'"`
	Ip             string    `json:"ip" xorm:"comment('IP地址') VARCHAR(64) 'ip'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
}

type SysLoginLog struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	UserName       string    `json:"user_name" xorm:"comment('用户名') VARCHAR(50) 'user_name'"`
	Status         string    `json:"status" xorm:"comment('登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）') VARCHAR(50) 'status'"`
	Ip             string    `json:"ip" xorm:"comment('IP地址') VARCHAR(64) 'ip'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
}

type SysMenu struct {
	Id             int       `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name           string    `json:"name" xorm:"comment('菜单名称') VARCHAR(50) 'name'"`
	ParentId       int       `json:"parent_id" xorm:"comment('父菜单ID，一级菜单为0') BIGINT(20) 'parent_id'"`
	Url            string    `json:"url" xorm:"comment('菜单URL,类型：1.普通页面（如用户管理， /sys/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)') VARCHAR(200) 'url'"`
	Perms          string    `json:"perms" xorm:"comment('授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)') VARCHAR(500) 'perms'"`
	Type           int64     `json:"type" xorm:"comment('类型   0：目录   1：菜单   2：按钮') INT(11) 'type'"`
	Icon           string    `json:"icon" xorm:"comment('菜单图标') VARCHAR(50) 'icon'"`
	OrderNum       int64     `json:"order_num" xorm:"comment('排序') INT(11) 'order_num'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysRole struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name           string    `json:"name" xorm:"comment('角色名称') VARCHAR(100) 'name'"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(100) 'remark'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysRoleDept struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	RoleId         int64     `json:"role_id" xorm:"comment('角色ID') BIGINT(20) 'role_id'"`
	DeptId         int64     `json:"dept_id" xorm:"comment('机构ID') BIGINT(20) 'dept_id'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
}

type SysRoleMenu struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	RoleId         int64     `json:"role_id" xorm:"comment('角色ID') BIGINT(20) 'role_id'"`
	MenuId         int64     `json:"menu_id" xorm:"comment('菜单ID') BIGINT(20) 'menu_id'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
}

type SysUser struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	Name           string    `json:"name" xorm:"not null comment('用户名') VARCHAR(50) 'name'"`
	NickName       string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(150) 'nick_name'"`
	Avatar         string    `json:"avatar" xorm:"comment('头像') VARCHAR(150) 'avatar'"`
	Password       string    `json:"password" xorm:"comment('密码') VARCHAR(100) 'password'"`
	Salt           string    `json:"salt" xorm:"comment('加密盐') VARCHAR(40) 'salt'"`
	Email          string    `json:"email" xorm:"comment('邮箱') VARCHAR(100) 'email'"`
	Mobile         string    `json:"mobile" xorm:"comment('手机号') VARCHAR(100) 'mobile'"`
	Status         int64     `json:"status" xorm:"comment('状态  0：禁用   1：正常') TINYINT(4) 'status'"`
	DeptId         int64     `json:"dept_id" xorm:"comment('机构ID') BIGINT(20) 'dept_id'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
	DelFlag        int64     `json:"del_flag" xorm:"default 0 comment('是否删除  -1：已删除  0：正常') TINYINT(4) 'del_flag'"`
}

type SysUserRole struct {
	Id             int64     `json:"id" xorm:"pk autoincr comment('编号') BIGINT(20) 'id'"`
	UserId         int64     `json:"user_id" xorm:"comment('用户ID') BIGINT(20) 'user_id'"`
	RoleId         int64     `json:"role_id" xorm:"comment('角色ID') BIGINT(20) 'role_id'"`
	CreateBy       string    `json:"create_by" xorm:"comment('创建人') VARCHAR(50) 'create_by'"`
	CreateTime     time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME 'create_time'"`
	LastUpdateBy   string    `json:"last_update_by" xorm:"comment('更新人') VARCHAR(50) 'last_update_by'"`
	LastUpdateTime time.Time `json:"last_update_time" xorm:"comment('更新时间') DATETIME 'last_update_time'"`
}
