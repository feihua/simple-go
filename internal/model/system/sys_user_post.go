package system

// SysUserPost 用户与岗位关联表
type SysUserPost struct {
	UserId int64 `gorm:"column:user_id;primary_key" json:"user_id"` // 用户id
	PostId int64 `gorm:"column:post_id;NOT NULL" json:"post_id"`    // 岗位id
}

func (m *SysUserPost) TableName() string {
	return "sys_user_post"
}
