package repositories

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/models"
)

func CreateSysLog(dto dto.LogDto) error {
	log := models.SysLog{}

	err := models.DB.Create(&log).Error

	return err
}

func GetSysLogList(current int, pageSize int) ([]models.SysLog, int) {

	var loginLog []models.SysLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.SysLog{}).Count(&total)
	return loginLog, total
}

func UpdateSysLog(logDto dto.LogDto) error {
	log := models.SysLog{}
	//if logDto.Username != "" {
	//	log.Username = logDto.Username
	//}
	//
	//if logDto.Password != "" {
	//	log.Password = logDto.Password
	//}

	err := models.DB.Model(&log).Update(&log).Error

	return err
}

func DeleteSysLogById(id int64) error {
	err := models.DB.Delete(&models.SysLog{Id: id}).Error
	return err
}

func CreateSysLoginLog(dto dto.LoginLogDto) error {
	log := models.SysLoginLog{}

	err := models.DB.Create(&log).Error

	return err
}

func GetSysLoginLogList(current int, pageSize int) ([]models.SysLoginLog, int) {
	var loginLog []models.SysLoginLog

	var total = 0
	models.DB.Limit(pageSize).Offset((current - 1) * pageSize).Find(&loginLog)

	models.DB.Model(&models.SysLoginLog{}).Count(&total)
	return loginLog, total
}

func UpdateSysLoginLog(loginLogDto dto.LoginLogDto) error {
	log := models.SysLoginLog{}

	err := models.DB.Model(&log).Update(&log).Error

	return err
}

func DeleteSysLoginLogById(id int64) error {
	err := models.DB.Delete(&models.SysLoginLog{Id: id}).Error
	return err
}
