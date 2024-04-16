package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// CreateLog 创建操作日志
func CreateLog(dto dto.LogDto) error {
	log := models.OperationLog{
		UserName:      dto.UserName,
		Operation:     dto.Operation,
		Method:        dto.Method,
		Params:        dto.Params,
		OperationTime: dto.OperationTime,
		Ip:            dto.Ip,
	}

	err := models.DB.Create(&log).Error

	return err
}

// QueryLogList 查询操作日志
func QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {

	var loginLog []models.OperationLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.OperationLog{}).Count(&total)
	return loginLog, total
}

// DeleteLogByIds 删除操作日志
func DeleteLogByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}

// CreateLoginLog 创建登录日志
func CreateLoginLog(dto dto.LoginLogDto) error {
	log := models.LoginLog{
		UserName: dto.UserName,
		Ip:       dto.Ip,
	}

	err := models.DB.Create(&log).Error

	return err
}

// QueryLoginLogList 查询登录日志
func QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	var loginLog []models.LoginLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.LoginLog{}).Count(&total)
	return loginLog, total
}

// DeleteLoginLogByIds 删除登录日志
func DeleteLoginLogByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}
