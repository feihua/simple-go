package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/notice"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
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

	req := rq.AddNoticeReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AddNoticeDto{
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      c.MustGet("userName").(string),
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

	req := rq.DeleteNoticeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
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

	req := rq.UpdateNoticeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateNoticeDto{
		Id:            req.Id,            // 公告ID
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		UpdateBy:      c.MustGet("userName").(string),
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

	req := rq.UpdateNoticeStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateNoticeStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
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
	req := rq.QueryNoticeDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryNoticeDetailDto{
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
	req := rq.QueryNoticeListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryNoticeListDto{
		PageNo:      req.PageNo,
		PageSize:    req.PageSize,
		NoticeTitle: req.NoticeTitle, // 公告标题
		NoticeType:  req.NoticeType,  // 公告类型（1:通知,2:公告）
		Status:      req.Status,      // 公告状态（0:关闭,1:正常 ）
	}
	list, total, err := r.Service.QueryNoticeList(item)
	if err != nil {
		result.FailWithMsg(c, result.NoticeError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}
