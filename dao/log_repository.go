package dao

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/jinzhu/gorm"
)

type LogDao struct {
	db *gorm.DB
}

func NewLogDao(DB *gorm.DB) *LogDao {
	return &LogDao{db: DB}
}

// CreateLog 创建操作日志
func (l LogDao) CreateLog(dto dto.LogDto) error {
	log := models.OperationLog{
		UserName:      dto.UserName,
		Operation:     dto.Operation,
		Method:        dto.Method,
		Params:        dto.Params,
		OperationTime: dto.OperationTime,
		Ip:            dto.Ip,
	}

	err := l.db.Create(&log).Error

	return err
}

// QueryLogList 查询操作日志
func (l LogDao) QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {

	var loginLog []models.OperationLog

	var total = 0
	l.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	l.db.Model(&models.OperationLog{}).Count(&total)
	return loginLog, total
}

// DeleteLogByIds 删除操作日志
func (l LogDao) DeleteLogByIds(ids []int64) error {
	return l.db.Where("id in (?)", ids).Delete(&models.OperationLog{}).Error
}

// CreateLoginLog 创建登录日志
func (l LogDao) CreateLoginLog(dto dto.LoginLogDto) error {
	log := models.LoginLog{
		UserName: dto.UserName,
		Ip:       dto.Ip,
	}

	err := l.db.Create(&log).Error

	return err
}

// QueryLoginLogList 查询登录日志
func (l LogDao) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	var loginLog []models.LoginLog

	var total = 0
	l.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	l.db.Model(&models.LoginLog{}).Count(&total)
	return loginLog, total
}

// DeleteLoginLogByIds 删除登录日志
func (l LogDao) DeleteLoginLogByIds(ids []int64) error {
	return l.db.Where("id in (?)", ids).Delete(&models.LoginLog{}).Error
}
