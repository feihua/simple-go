package log

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogServiceImpl struct {
}

func (l *LogServiceImpl) CreateLog(dto dto.LogDto) error {
	return dao.CreateSysLog(dto)
}

func (l *LogServiceImpl) QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {
	return dao.QueryLogList(current, pageSize)
}

func (l *LogServiceImpl) DeleteLogByIds(ids []int64) error {
	return dao.DeleteSysLogByIds(ids)
}

func (l *LogServiceImpl) CreateLoginLog(dto dto.LoginLogDto) error {
	return dao.CreateSysLoginLog(dto)
}

func (l *LogServiceImpl) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	return dao.QueryLoginLogList(current, pageSize)
}

func (l *LogServiceImpl) DeleteLoginLogByIds(ids []int64) error {
	return dao.DeleteSysLoginLogByIds(ids)
}
