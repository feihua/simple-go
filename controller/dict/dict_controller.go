package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/pkg/result"
	"github.com/feihua/simple-go/services/dict"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

// DictController 字典相关
/*
Author: LiuFeiHua
Date: 2024/4/15 16:48
*/
type DictController struct {
}

func NewDictController() *DictController {
	return &DictController{}
}

// CreateDict 创建字典
func (d DictController) CreateDict(c *gin.Context) {

	req := requests.DictRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service = &dict.DictServiceImpl{}

	dictDto := dto.DictDto{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Remarks:     req.Remarks,
	}

	err = service.CreateDict(dictDto)
	if err != nil {
		result.FailWithMsg(c, result.DictError, err.Error())
	} else {
		result.Ok(c)
	}
}

// QueryDictList 查询字典列表
func (d DictController) QueryDictList(c *gin.Context) {
	req := requests.DictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service = &dict.DictServiceImpl{}

	dictList, total := service.QueryDictList(pageNum, size)

	result.OkWithData(c, gin.H{"list": dictList, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

// UpdateDict 更新字典
func (d DictController) UpdateDict(c *gin.Context) {

	req := requests.DictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service = &dict.DictServiceImpl{}

	dictDto := dto.DictDto{
		Id:          req.Id,
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Remarks:     req.Remarks,
	}
	err = service.UpdateDict(dictDto)
	if err != nil {
		result.FailWithMsg(c, result.DictError, err.Error())
	} else {
		result.Ok(c)
	}
}

// DeleteDictByIds 删除字典
func (d DictController) DeleteDictByIds(c *gin.Context) {

	req := requests.DeleteDictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		result.Fail(c, result.ParamsError)
		return
	}

	var service = &dict.DictServiceImpl{}

	err = service.DeleteDictByIds(req.Ids)
	if err != nil {
		result.FailWithMsg(c, result.DictError, err.Error())
	} else {
		result.Ok(c)
	}
}
