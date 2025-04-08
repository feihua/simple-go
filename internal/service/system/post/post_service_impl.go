package post

import (
	"github.com/feihua/simple-go/internal/dao/system"
	a "github.com/feihua/simple-go/internal/dto/system"
	b "github.com/feihua/simple-go/internal/model/system"
)

// PostServiceImpl 岗位信息操作实现
type PostServiceImpl struct {
	Dao *system.PostDao
}

func NewPostServiceImpl(dao *system.PostDao) PostService {
	return &PostServiceImpl{
		Dao: dao,
	}
}

// CreatePost 添加岗位信息
func (s *PostServiceImpl) CreatePost(dto a.AddPostDto) error {
	return s.Dao.CreatePost(dto)
}

// DeletePostByIds 删除岗位信息
func (s *PostServiceImpl) DeletePostByIds(ids []int64) error {
	return s.Dao.DeletePostByIds(ids)
}

// UpdatePost 更新岗位信息
func (s *PostServiceImpl) UpdatePost(dto a.UpdatePostDto) error {
	return s.Dao.UpdatePost(dto)
}

// UpdatePostStatus 更新岗位信息状态
func (s *PostServiceImpl) UpdatePostStatus(dto a.UpdatePostStatusDto) error {
	return s.Dao.UpdatePostStatus(dto)
}

// QueryPostDetail 查询岗位信息详情
func (s *PostServiceImpl) QueryPostDetail(dto a.QueryPostDetailDto) (b.Post, error) {
	return s.Dao.QueryPostDetail(dto)
}

// QueryPostList 查询岗位信息列表
func (s *PostServiceImpl) QueryPostList(dto a.QueryPostListDto) ([]b.Post, int64) {
	return s.Dao.QueryPostList(dto)
}
