package log

import (
	"github.com/feihua/simple-go/dao"
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogService struct {
}

func (l *LogService) CreateLog(dto dto.LogDto) error {
	return dao.CreateSysLog(dto)
}

func (l *LogService) QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {
	return dao.QueryLogList(current, pageSize)
}

func (l *LogService) DeleteLogById(id int64) error {
	return dao.DeleteSysLogById(id)
}

func (l *LogService) CreateLoginLog(dto dto.LoginLogDto) error {
	return dao.CreateSysLoginLog(dto)
}

func (l *LogService) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	return dao.QueryLoginLogList(current, pageSize)
}

func (l *LogService) DeleteLoginLogById(id int64) error {
	return dao.DeleteSysLoginLogById(id)
}
