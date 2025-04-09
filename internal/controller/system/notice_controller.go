package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/notice"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// NoticeController 通知公告相关
type NoticeController struct {
	Service notice.NoticeService
}

func NewNoticeController(Service notice.NoticeService) *NoticeController {
	return &NoticeController{Service: Service}
}

// CreateNotice 添加通知公告
func (r NoticeController) CreateNotice(c *gin.Context) {

	req := b.AddNoticeReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.AddNoticeDto{
		Id:            req.Id,            // 公告ID
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      req.CreateBy,      // 创建者
		CreateTime:    req.CreateTime,    // 创建时间
		UpdateBy:      req.UpdateBy,      // 更新者
		UpdateTime:    req.UpdateTime,    // 更新时间
	}

	err = r.Service.CreateNotice(item)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteNoticeByIds 删除通知公告
func (r NoticeController) DeleteNoticeByIds(c *gin.Context) {

	req := b.DeleteNoticeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	err = r.Service.DeleteNoticeByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateNotice 更新通知公告
func (r NoticeController) UpdateNotice(c *gin.Context) {

	req := b.UpdateNoticeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateNoticeDto{
		Id:            req.Id,            // 公告ID
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      req.CreateBy,      // 创建者
		CreateTime:    req.CreateTime,    // 创建时间
		UpdateBy:      req.UpdateBy,      // 更新者
		UpdateTime:    req.UpdateTime,    // 更新时间
	}
	err = r.Service.UpdateNotice(item)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateNoticeStatus 更新通知公告状态
func (r NoticeController) UpdateNoticeStatus(c *gin.Context) {

	req := b.UpdateNoticeStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateNoticeStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateNoticeStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryNoticeDetail 查询通知公告详情
func (r NoticeController) QueryNoticeDetail(c *gin.Context) {
	req := b.QueryNoticeDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryNoticeDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryNoticeDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryNoticeList 查询通知公告列表
func (r NoticeController) QueryNoticeList(c *gin.Context) {
	req := b.QueryNoticeListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryNoticeListDto{
		PageNo:        req.PageNo,
		PageSize:      req.PageSize,
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
	}
	list, total := r.Service.QueryNoticeList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
