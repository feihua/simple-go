package log

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

// LogServiceImpl 日志操作实现
/*
Author: LiuFeiHua
Date: 2024/4/16 11:23
*/
type LogServiceImpl struct {
}

// CreateLog 创建操作日志
func (l *LogServiceImpl) CreateLog(dto dto.LogDto) error {
	return dao.CreateLog(dto)
}

// QueryLogList 查询操作日志
func (l *LogServiceImpl) QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {
	return dao.QueryLogList(current, pageSize)
}

// DeleteLogByIds 删除操作日志
func (l *LogServiceImpl) DeleteLogByIds(ids []int64) error {
	return dao.DeleteLogByIds(ids)
}

// CreateLoginLog 创建登录日志
func (l *LogServiceImpl) CreateLoginLog(dto dto.LoginLogDto) error {
	return dao.CreateLoginLog(dto)
}

// QueryLoginLogList 查询登录日志
func (l *LogServiceImpl) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	return dao.QueryLoginLogList(current, pageSize)
}

// DeleteLoginLogByIds 删除登录日志
func (l *LogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return dao.DeleteLoginLogByIds(ids)
}
