package dao

import (
	"github.com/feihua/simple-go/internal/dto"
	"github.com/feihua/simple-go/internal/model"
	"gorm.io/gorm"
)

type LogDao struct {
	db *gorm.DB
}

func NewLogDao(DB *gorm.DB) *LogDao {
	return &LogDao{db: DB}
}

// CreateLog 创建操作日志
func (l LogDao) CreateLog(dto dto.LogDto) error {
	log := model.OperateLog{
		UserName:       dto.UserName,
		Operation:      dto.Operation,
		Method:         dto.Method,
		RequestParams:  dto.RequestParams,
		ResponseParams: dto.ResponseParams,
		Time:           dto.OperationTime,
		Ip:             dto.Ip,
	}

	err := l.db.Create(&log).Error

	return err
}

// QueryLogList 查询操作日志
func (l LogDao) QueryLogList(current int, pageSize int) ([]model.OperateLog, int64) {

	var loginLog []model.OperateLog

	var total int64 = 0
	l.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	l.db.Model(&model.OperateLog{}).Count(&total)
	return loginLog, total
}

// DeleteLogByIds 删除操作日志
func (l LogDao) DeleteLogByIds(ids []int64) error {
	return l.db.Where("id in (?)", ids).Delete(&model.OperateLog{}).Error
}

// CreateLoginLog 创建登录日志
func (l LogDao) CreateLoginLog(dto dto.LoginLogDto) error {
	log := model.LoginLog{
		UserName: dto.UserName,
		Ip:       dto.Ip,
	}

	err := l.db.Create(&log).Error

	return err
}

// QueryLoginLogList 查询登录日志
func (l LogDao) QueryLoginLogList(current int, pageSize int) ([]model.LoginLog, int64) {
	var loginLog []model.LoginLog

	var total int64 = 0
	l.db.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	l.db.Model(&model.LoginLog{}).Count(&total)
	return loginLog, total
}

// DeleteLoginLogByIds 删除登录日志
func (l LogDao) DeleteLoginLogByIds(ids []int64) error {
	return l.db.Where("id in (?)", ids).Delete(&model.LoginLog{}).Error
}
