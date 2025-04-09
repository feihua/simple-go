package system

// AddUserPostDto 添加用户与岗位关联请求参数
type AddUserPostDto struct {
	UserId int64 `json:"userId"` // 用户id
	PostId int64 `json:"postId"` // 岗位id
}

// DeleteUserPostDto 删除用户与岗位关联请求参数
type DeleteUserPostDto struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserPostDto 修改用户与岗位关联请求参数
type UpdateUserPostDto struct {
	UserId int64 `json:"userId"` // 用户id
	PostId int64 `json:"postId"` // 岗位id
}

// UpdateUserPostStatusDto 修改用户与岗位关联状态请求参数
type UpdateUserPostStatusDto struct {
	Ids    []int64 `json:"ids"`    // id
	Status int32   `json:"status"` // 状态（0:关闭,1:正常 ）
}

// QueryUserPostDetailDto 查询用户与岗位关联详情请求参数
type QueryUserPostDetailDto struct {
	Id int64 `json:"id"` // id
}

// QueryUserPostListDto 查询用户与岗位关联列表请求参数
type QueryUserPostListDto struct {
	PageNo   int   `json:"pageNo" default:"1"`    // 第几页
	PageSize int   `json:"pageSize" default:"10"` // 每页的数量
	UserId   int64 `json:"userId"`                // 用户id
	PostId   int64 `json:"postId"`                // 岗位id
}
