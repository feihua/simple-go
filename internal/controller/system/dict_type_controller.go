package system

import (
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/dict_type"
	rq "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// DictTypeController 字典类型相关
type DictTypeController struct {
	Service dict_type.DictTypeService
}

func NewDictTypeController(Service dict_type.DictTypeService) *DictTypeController {
	return &DictTypeController{Service: Service}
}

// CreateDictType 添加字典类型
func (r DictTypeController) CreateDictType(c *gin.Context) {

	req := rq.AddDictTypeReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.AddDictTypeDto{
		DictName: req.DictName, // 字典名称
		DictType: req.DictType, // 字典类型
		Status:   req.Status,   // 状态（0：停用，1:正常）
		Remark:   req.Remark,   // 备注
		CreateBy: c.MustGet("userName").(string),
	}

	err = r.Service.CreateDictType(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteDictTypeByIds 删除字典类型
func (r DictTypeController) DeleteDictTypeByIds(c *gin.Context) {

	req := rq.DeleteDictTypeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteDictTypeByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDictType 更新字典类型
func (r DictTypeController) UpdateDictType(c *gin.Context) {

	req := rq.UpdateDictTypeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateDictTypeDto{
		Id:       req.Id,       // 字典主键
		DictName: req.DictName, // 字典名称
		DictType: req.DictType, // 字典类型
		Status:   req.Status,   // 状态（0：停用，1:正常）
		Remark:   req.Remark,   // 备注
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateDictType(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDictTypeStatus 更新字典类型状态
func (r DictTypeController) UpdateDictTypeStatus(c *gin.Context) {

	req := rq.UpdateDictTypeStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.UpdateDictTypeStatusDto{
		Ids:      req.Ids,
		Status:   req.Status,
		UpdateBy: c.MustGet("userName").(string),
	}
	err = r.Service.UpdateDictTypeStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryDictTypeDetail 查询字典类型详情
func (r DictTypeController) QueryDictTypeDetail(c *gin.Context) {
	req := rq.QueryDictTypeDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryDictTypeDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryDictTypeDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.OkWithData(c, data)
	}
}

// QueryDictTypeList 查询字典类型列表
func (r DictTypeController) QueryDictTypeList(c *gin.Context) {
	req := rq.QueryDictTypeListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := d.QueryDictTypeListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		DictName: req.DictName, // 字典名称
		DictType: req.DictType, // 字典类型
		Status:   req.Status,   // 状态（0：停用，1:正常）
	}
	list, total, err := r.Service.QueryDictTypeList(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
	}
}
