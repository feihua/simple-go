package system

// UserPost 用户与岗位关联
type UserPost struct {
	UserId int64 `json:"userId"` // 用户id
	PostId int64 `json:"postId"` // 岗位id
}
