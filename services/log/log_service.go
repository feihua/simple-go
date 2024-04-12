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

func (l *LogService) GetLogList(current int, pageSize int) ([]models.SysLog, int) {
	return repositories.GetSysLogList(current, pageSize)
}

func (l *LogService) UpdateLog(logDto dto.LogDto) error {
	return repositories.UpdateSysLog(logDto)
}

func (l *LogService) DeleteLogById(id int64) error {
	return repositories.DeleteSysLogById(id)
}

func (l *LogService) CreateLoginLog(dto dto.LoginLogDto) error {
	return repositories.CreateSysLoginLog(dto)
}

func (l *LogService) GetLoginLogList(current int, pageSize int) ([]models.SysLoginLog, int) {
	return repositories.GetSysLoginLogList(current, pageSize)
}

func (l *LogService) UpdateLoginLog(logDto dto.LoginLogDto) error {
	return repositories.UpdateSysLoginLog(logDto)
}

func (l *LogService) DeleteLoginLogById(id int64) error {
	return repositories.DeleteSysLoginLogById(id)
}
