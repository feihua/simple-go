package system

import "time"

// AddUserPostReqVo 添加用户与岗位关联请求参数
type AddUserPostReqVo struct {
	UserId int64 `json:"userId"` //用户id
	PostId int64 `json:"postId"` //岗位id
}

// DeleteUserPostReqVo 删除用户与岗位关联请求参数
type DeleteUserPostReqVo struct {
	Ids []int64 `json:"ids"`
}

// UpdateUserPostReqVo 修改用户与岗位关联请求参数
type UpdateUserPostReqVo struct {
	UserId int64 `json:"userId"` //用户id
	PostId int64 `json:"postId"` //岗位id
}

// UpdateUserPostStatusReqVo 修改用户与岗位关联状态请求参数
type UpdateUserPostStatusReqVo struct {
	Ids    []int64 `json:"ids"`    //id
	Status int32   `json:"status"` //状态（0:关闭,1:正常 ）
}

// QueryUserPostDetailReqVo 查询用户与岗位关联详情请求参数
type QueryUserPostDetailReqVo struct {
	Id int64 `json:"id"` //id
}

// QueryUserPostListReqVo 查询用户与岗位关联列表请求参数
type QueryUserPostListReqVo struct {
	PageNo   int   `json:"pageNo" default:"1"`    //第几页
	PageSize int   `json:"pageSize" default:"10"` //每页的数量
	UserId   int64 `json:"userId"`                //用户id
	PostId   int64 `json:"postId"`                //岗位id
}
