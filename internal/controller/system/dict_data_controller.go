package system

import (
	a "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/internal/service/system/dict_data"
	b "github.com/feihua/simple-go/internal/vo/system/req"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/gin-gonic/gin"
)

// DictDataController 字典数据相关
type DictDataController struct {
	Service dict_data.DictDataService
}

func NewDictDataController(Service dict_data.DictDataService) *DictDataController {
	return &DictDataController{Service: Service}
}

// CreateDictData 添加字典数据
func (r DictDataController) CreateDictData(c *gin.Context) {

	req := b.AddDictDataReqVo{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.AddDictDataDto{
		Id:         req.Id,         // 字典编码
		DictSort:   req.DictSort,   // 字典排序
		DictLabel:  req.DictLabel,  // 字典标签
		DictValue:  req.DictValue,  // 字典键值
		DictType:   req.DictType,   // 字典类型
		CssClass:   req.CssClass,   // 样式属性（其他样式扩展）
		ListClass:  req.ListClass,  // 表格回显样式
		IsDefault:  req.IsDefault,  // 是否默认（Y是 N否）
		Status:     req.Status,     // 状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}

	err = r.Service.CreateDictData(item)
	if err != nil {
		result.FailWithMsg(c, result.DictDataError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteDictDataByIds 删除字典数据
func (r DictDataController) DeleteDictDataByIds(c *gin.Context) {

	req := b.DeleteDictDataReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	err = r.Service.DeleteDictDataByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.DictDataError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDictData 更新字典数据
func (r DictDataController) UpdateDictData(c *gin.Context) {

	req := b.UpdateDictDataReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateDictDataDto{
		Id:         req.Id,         // 字典编码
		DictSort:   req.DictSort,   // 字典排序
		DictLabel:  req.DictLabel,  // 字典标签
		DictValue:  req.DictValue,  // 字典键值
		DictType:   req.DictType,   // 字典类型
		CssClass:   req.CssClass,   // 样式属性（其他样式扩展）
		ListClass:  req.ListClass,  // 表格回显样式
		IsDefault:  req.IsDefault,  // 是否默认（Y是 N否）
		Status:     req.Status,     // 状态（0：停用，1:正常）
		Remark:     req.Remark,     // 备注
		CreateBy:   req.CreateBy,   // 创建者
		CreateTime: req.CreateTime, // 创建时间
		UpdateBy:   req.UpdateBy,   // 更新者
		UpdateTime: req.UpdateTime, // 更新时间
	}
	err = r.Service.UpdateDictData(item)
	if err != nil {
		result.FailWithMsg(c, result.DictDataError, err.Error())
	} else {
		result.Ok(c)
	}
}

// UpdateDictDataStatus 更新字典数据状态
func (r DictDataController) UpdateDictDataStatus(c *gin.Context) {

	req := b.UpdateDictDataStatusReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.UpdateDictDataStatusDto{
		Ids:    req.Ids,
		Status: req.Status,
	}
	err = r.Service.UpdateDictDataStatus(item)
	if err != nil {
		result.FailWithMsg(c, result.DictDataError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryDictDataDetail 查询字典数据详情
func (r DictDataController) QueryDictDataDetail(c *gin.Context) {
	req := b.QueryDictDataDetailReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryDictDataDetailDto{
		Id: req.Id,
	}
	data, err := r.Service.QueryDictDataDetail(item)
	if err != nil {
		result.FailWithMsg(c, result.DictDataError, err.Error())
	} else {
		result.OkWithData(c, gin.H{"data": data})
	}
}

// QueryDictDataList 查询字典数据列表
func (r DictDataController) QueryDictDataList(c *gin.Context) {
	req := b.QueryDictDataListReqVo{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMsg(c, result.ParamsError, err.Error())
		return
	}

	item := a.QueryDictDataListDto{
		PageNo:    req.PageNo,
		PageSize:  req.PageSize,
		DictLabel: req.DictLabel, // 字典标签
		DictValue: req.DictValue, // 字典键值
		DictType:  req.DictType,  // 字典类型
		CssClass:  req.CssClass,  // 样式属性（其他样式扩展）
		ListClass: req.ListClass, // 表格回显样式
		IsDefault: req.IsDefault, // 是否默认（Y是 N否）
		Status:    req.Status,    // 状态（0：停用，1:正常）
	}
	list, total := r.Service.QueryDictDataList(item)
	result.OkWithData(c, gin.H{"list": list, "success": true, "current": req.PageNo, "total": total, "pageSize": req.PageSize})
}
