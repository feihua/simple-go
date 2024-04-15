package dict

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/dict"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (d DictController) CreateDict(c *gin.Context) {

	req := requests.DictRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictService = &dict.DictServiceImpl{}

	u := dto.DictDto{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Remarks:     req.Remarks,
	}
	result := service.CreateDict(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (d DictController) QueryDictList(c *gin.Context) {
	req := requests.DictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service dict.DictService = &dict.DictServiceImpl{}

	result, total := service.QueryDictList(pageNum, size)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func (d DictController) UpdateDict(c *gin.Context) {

	req := requests.DictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictService = &dict.DictServiceImpl{}

	u := dto.DictDto{
		Id:          req.Id,
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Remarks:     req.Remarks,
	}
	result := service.UpdateDict(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (d DictController) DeleteDictByIds(c *gin.Context) {

	req := requests.DeleteDictRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service dict.DictService = &dict.DictServiceImpl{}

	result := service.DeleteDictByIds(req.Ids)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
