package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/post"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// PostController 岗位信息相关
type PostController struct {
	Service post.PostService
}

func NewPostController(Service post.PostService) *PostController {
	return &PostController{Service: Service}
}

// CreatePost 添加岗位信息
func (r PostController) CreatePost(c *gin.Context) {

	req := b.AddPostReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddPostDto{
		Id:         req.Id,         // 岗位id
		PostCode:   req.PostCode,   // 岗位编码
		PostName:   req.PostName,   // 岗位名称
		Sort:       req.Sort,       // 显示顺序
		Status:     req.Status,     // 岗位状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}

	err = r.Service.CreatePost(item)
	if err != nil {
		result.FailWithMsg(c, result.PostError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeletePostByIds 删除岗位信息
func (r PostController) DeletePostByIds(c *gin.Context) {

	req := b.DeletePostReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeletePostByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.PostError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdatePost 更新岗位信息
func (r PostController) UpdatePost(c *gin.Context) {

	req := b.UpdatePostReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdatePostDto{
		Id:         req.Id,         // 岗位id
		PostCode:   req.PostCode,   // 岗位编码
		PostName:   req.PostName,   // 岗位名称
		Sort:       req.Sort,       // 显示顺序
		Status:     req.Status,     // 岗位状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}
	err = r.Service.UpdatePost(item)
	if err != nil {
		result.FailWithMsg(c, result.PostError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdatePostStatus 更新岗位信息状态
func (r PostController) UpdatePostStatus(c *gin.Context) {

	req := b.UpdatePostStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdatePostStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdatePostStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.PostError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryPostDetail 查询岗位信息详情
func (r PostController) QueryPostDetail(c *gin.Context) {
	req := b.QueryPostDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryPostDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryPostDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.PostError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryPostList 查询岗位信息列表
func (r PostController) QueryPostList(c *gin.Context) {
	req := b.QueryPostListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryPostListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		PostCode: req.PostCode, // 岗位编码
		PostName: req.PostName, // 岗位名称
		Status:   req.Status,   // 岗位状态（0：停用，1:正常）
	}
	list, total := r.Service.QueryPostList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
