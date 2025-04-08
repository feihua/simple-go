package post

import (
	b "github.com/feihua/simple-go/internal/dto/system"
	a "github.com/feihua/simple-go/internal/model/system"
)

// PostService 岗位信息操作接口
type PostService interface {
	// CreatePost 添加岗位信息
	CreatePost(dto b.AddPostDto) error
	// DeletePostByIds 删除岗位信息
	DeletePostByIds(ids []int64) error
	// UpdatePost 更新岗位信息
	UpdatePost(dto b.UpdatePostDto) error
	// UpdatePostStatus 更新岗位信息状态
	UpdatePostStatus(dto b.UpdatePostStatusDto) error
	// QueryPostDetail 查询岗位信息详情
	QueryPostDetail(dto b.QueryPostDetailDto) (a.Post, error)
	// QueryPostList 查询岗位信息列表
	QueryPostList(dto b.QueryPostListDto) ([]a.Post, int64)
}
