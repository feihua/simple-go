package system

// UserPost 用户与岗位关联
type UserPost struct {
	UserId int64 `gorm:"column:user_id" json:"userId"` // 用户id
	PostId int64 `gorm:"column:post_id" json:"postId"` // 岗位id
}

func (model UserPost) TableName() string {
	return "sys_user_post"
}
