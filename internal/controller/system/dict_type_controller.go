package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/dict_type"
	b "github.com/feihua/simple-go/internal/vo/system/req"
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

	req := b.AddDictTypeReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.AddDictTypeDto{
		Id:         req.Id,         // 字典主键
		DictName:   req.DictName,   // 字典名称
		DictType:   req.DictType,   // 字典类型
		Status:     req.Status,     // 状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
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

	req := b.DeleteDictTypeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
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

	req := b.UpdateDictTypeReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateDictTypeDto{
		Id:         req.Id,         // 字典主键
		DictName:   req.DictName,   // 字典名称
		DictType:   req.DictType,   // 字典类型
		Status:     req.Status,     // 状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
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

	req := b.UpdateDictTypeStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.UpdateDictTypeStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
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
	req := b.QueryDictTypeDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryDictTypeDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryDictTypeDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.DictTypeError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryDictTypeList 查询字典类型列表
func (r DictTypeController) QueryDictTypeList(c *gin.Context) {
	req := b.QueryDictTypeListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	item := a.QueryDictTypeListDto{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		DictName: req.DictName, // 字典名称
		DictType: req.DictType, // 字典类型
		Status:   req.Status,   // 状态（0：停用，1:正常）
	}
	list, total := r.Service.QueryDictTypeList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
