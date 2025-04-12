package post

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
	"time"
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
func (s *PostServiceImpl) CreatePost(dto d.AddPostDto) error {
	post, err := s.Dao.QueryPostByName(dto.PostName)
	if err != nil {
		return err
	}
	if post != nil {
		return errors.New("岗位名称已存在")
	}

	code, err := s.Dao.QueryPostByCode(dto.PostCode)
	if err != nil {
		return err
	}
	if code != nil {
		return errors.New("岗位编码已存在")
	}
	return s.Dao.CreatePost(dto)
}

// DeletePostByIds 删除岗位信息
func (s *PostServiceImpl) DeletePostByIds(ids []int64) error {
	return s.Dao.DeletePostByIds(ids)
}

// UpdatePost 更新岗位信息
func (s *PostServiceImpl) UpdatePost(dto d.UpdatePostDto) error {
	item, err := s.Dao.QueryPostById(dto.Id)

	if err != nil {
		return err
	}

	if item == nil {
		return errors.New("岗位信息不存在")
	}

	post, err := s.Dao.QueryPostByName(dto.PostName)
	if err != nil {
		return err
	}
	if post != nil && item.Id != dto.Id {
		return errors.New("岗位名称已存在")
	}

	code, err := s.Dao.QueryPostByCode(dto.PostCode)
	if err != nil {
		return err
	}
	if code != nil && item.Id != dto.Id {
		return errors.New("岗位编码已存在")
	}

	dto.CreateBy = item.CreateBy
	dto.CreateTime = item.CreateTime
	dto.UpdateTime = time.Now()
	return s.Dao.UpdatePost(dto)
}

// UpdatePostStatus 更新岗位信息状态
func (s *PostServiceImpl) UpdatePostStatus(dto d.UpdatePostStatusDto) error {
	return s.Dao.UpdatePostStatus(dto)
}

// QueryPostDetail 查询岗位信息详情
func (s *PostServiceImpl) QueryPostDetail(dto d.QueryPostDetailDto) (*d.QueryPostListDtoResp, error) {
	item, err := s.Dao.QueryPostDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("岗位信息不存在")
	}

	return &d.QueryPostListDtoResp{
		Id:         item.Id,                             // 岗位id
		PostCode:   item.PostCode,                       // 岗位编码
		PostName:   item.PostName,                       // 岗位名称
		Sort:       item.Sort,                           // 显示顺序
		Status:     item.Status,                         // 岗位状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryPostById 根据id查询岗位信息详情
func (s *PostServiceImpl) QueryPostById(id int64) (*d.QueryPostListDtoResp, error) {
	item, err := s.Dao.QueryPostById(id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("岗位信息不存在")
	}

	return &d.QueryPostListDtoResp{
		Id:         item.Id,                             // 岗位id
		PostCode:   item.PostCode,                       // 岗位编码
		PostName:   item.PostName,                       // 岗位名称
		Sort:       item.Sort,                           // 显示顺序
		Status:     item.Status,                         // 岗位状态（0：停用，1:正常）
		Remark:     item.Remark,                         // 备注
		CreateBy:   item.CreateBy,                       // 创建者
		CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                       // 更新者
		UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
	}, nil

}

// QueryPostList 查询岗位信息列表
func (s *PostServiceImpl) QueryPostList(dto d.QueryPostListDto) ([]*d.QueryPostListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryPostList(dto)

	if err != nil {
		return nil, 0, err
	}

	list := make([]*d.QueryPostListDtoResp, 0)

	for _, item := range result {
		resp := &d.QueryPostListDtoResp{
			Id:         item.Id,                             // 岗位id
			PostCode:   item.PostCode,                       // 岗位编码
			PostName:   item.PostName,                       // 岗位名称
			Sort:       item.Sort,                           // 显示顺序
			Status:     item.Status,                         // 岗位状态（0：停用，1:正常）
			Remark:     item.Remark,                         // 备注
			CreateBy:   item.CreateBy,                       // 创建者
			CreateTime: utils.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:   item.UpdateBy,                       // 更新者
			UpdateTime: utils.TimeToString(item.UpdateTime), // 更新时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
