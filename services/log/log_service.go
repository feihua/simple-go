package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
	"github.com/feihua/simple-go/repositories"
)

type LogService struct {
}

func (l *LogService) CreateLog(dto dto.LogDto) error {
	return repositories.CreateSysLog(dto)
}

func (l *LogService) QueryLogList(current int, pageSize int) ([]models.OperationLog, int) {
	return repositories.QueryLogList(current, pageSize)
}

func (l *LogService) DeleteLogById(id int64) error {
	return repositories.DeleteSysLogById(id)
}

func (l *LogService) CreateLoginLog(dto dto.LoginLogDto) error {
	return repositories.CreateSysLoginLog(dto)
}

func (l *LogService) QueryLoginLogList(current int, pageSize int) ([]models.LoginLog, int) {
	return repositories.QueryLoginLogList(current, pageSize)
}

func (l *LogService) DeleteLoginLogById(id int64) error {
	return repositories.DeleteSysLoginLogById(id)
}
