package post

import (
	d "github.com/feihua/simple-go/internal/dto/system"
)

// PostService 岗位信息操作接口
type PostService interface {
	// CreatePost 添加岗位信息
	CreatePost(dto d.AddPostDto) error
	// DeletePostByIds 删除岗位信息
	DeletePostByIds(ids []int64) error
	// UpdatePost 更新岗位信息
	UpdatePost(dto d.UpdatePostDto) error
	// UpdatePostStatus 更新岗位信息状态
	UpdatePostStatus(dto d.UpdatePostStatusDto) error
	// QueryPostDetail 查询岗位信息详情
	QueryPostDetail(dto d.QueryPostDetailDto) (*d.QueryPostListDtoResp, error)
	// QueryPostById 根据id查询岗位信息详情
	QueryPostById(id int64) (*d.QueryPostListDtoResp, error)
	// QueryPostList 查询岗位信息列表
	QueryPostList(dto d.QueryPostListDto) ([]*d.QueryPostListDtoResp, int64, error)
}
