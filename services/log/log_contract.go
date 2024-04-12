package log

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

type LogContract interface {
	CreateLog(dto dto.LogDto) error

	GetLogList(current int, pageSize int) ([]models.SysLog, int)

	UpdateLog(logDto dto.LogDto) error

	DeleteLogById(id int64) error

	CreateLoginLog(dto dto.LoginLogDto) error

	GetLoginLogList(current int, pageSize int) ([]models.SysLoginLog, int)

	UpdateLoginLog(logDto dto.LoginLogDto) error

	DeleteLoginLogById(id int64) error
}
