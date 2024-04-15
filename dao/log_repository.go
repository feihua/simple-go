package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateSysLog(dto dto.LogDto) error {
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

func QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {

	var loginLog []models.OperationLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.OperationLog{}).Count(&total)
	return loginLog, total
}

func DeleteSysLogByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}

func CreateSysLoginLog(dto dto.LoginLogDto) error {
	log := models.LoginLog{
		UserName: dto.UserName,
		Ip:       dto.Ip,
	}

	err := models.DB.Create(&log).Error

	return err
}

func QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	var loginLog []models.LoginLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.LoginLog{}).Count(&total)
	return loginLog, total
}

func DeleteSysLoginLogByIds(ids []int64) error {
	return models.DB.Where("id in (?)", ids).Delete(&models.Dept{}).Error
}
