package system

import (
	"errors"
	"github.com/feihua/simple-go/internal/dto/system"
	m "github.com/feihua/simple-go/internal/model/system"
	"gorm.io/gorm"
	"time"
)

type OperateLogDao struct {
	db *gorm.DB
}

func NewOperateLogDao(DB *gorm.DB) *OperateLogDao {
	return &OperateLogDao{
		db: DB,
	}
}

// CreateOperateLog 添加操作日志记录
func (b OperateLogDao) CreateOperateLog(dto system.AddOperateLogDto) error {
	item := m.OperateLog{
		Title:           dto.Title,           // 模块标题
		BusinessType:    dto.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          dto.Method,          // 方法名称
		RequestMethod:   dto.RequestMethod,   // 请求方式
		OperatorType:    dto.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     dto.OperateName,     // 操作人员
		DeptName:        dto.DeptName,        // 部门名称
		OperateUrl:      dto.OperateUrl,      // 请求URL
		OperateIp:       dto.OperateIp,       // 主机地址
		OperateLocation: dto.OperateLocation, // 操作地点
		OperateParam:    dto.OperateParam,    // 请求参数
		JsonResult:      dto.JsonResult,      // 返回参数
		Platform:        dto.Platform,        // 平台信息
		Browser:         dto.Browser,         // 浏览器类型
		Version:         dto.Version,         // 浏览器版本
		Os:              dto.Os,              // 操作系统
		Arch:            dto.Arch,            // 体系结构信息
		Engine:          dto.Engine,          // 渲染引擎信息
		EngineDetails:   dto.EngineDetails,   // 渲染引擎详细信息
		Extra:           dto.Extra,           // 其他信息（可选）
		Status:          dto.Status,          // 操作状态(0:异常,正常)
		ErrorMsg:        dto.ErrorMsg,        // 错误消息
		OperateTime:     time.Now(),          // 操作时间
		CostTime:        dto.CostTime,        // 消耗时间
	}

	return b.db.Create(&item).Error
}

// DeleteOperateLogByIds 根据id删除操作日志记录
func (b OperateLogDao) DeleteOperateLogByIds(ids []int64) error {
	return b.db.Where("id in (?)", ids).Delete(&m.OperateLog{}).Error
}

// QueryOperateLogDetail 查询操作日志记录详情
func (b OperateLogDao) QueryOperateLogDetail(dto system.QueryOperateLogDetailDto) (*m.OperateLog, error) {
	var item m.OperateLog
	err := b.db.Where("id", dto.Id).First(&item).Error

	switch {
	case err == nil:
		return &item, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	default:
		return nil, err
	}
}

// QueryOperateLogList 查询操作日志记录列表
func (b OperateLogDao) QueryOperateLogList(dto system.QueryOperateLogListDto) ([]*m.OperateLog, int64, error) {
	pageNo := dto.PageNo
	pageSize := dto.PageSize

	var total int64 = 0
	var list []*m.OperateLog
	tx := b.db.Model(&m.OperateLog{})
	if len(dto.Title) > 0 {
		tx.Where("title like %?%", dto.Title) // 模块标题
	}
	if dto.BusinessType != nil {
		tx.Where("business_type=?", dto.BusinessType) // 业务类型（0其它 1新增 2修改 3删除）
	}
	if len(dto.Method) > 0 {
		tx.Where("method like %?%", dto.Method) // 方法名称
	}
	if len(dto.RequestMethod) > 0 {
		tx.Where("request_method like %?%", dto.RequestMethod) // 请求方式
	}
	if dto.OperatorType != nil {
		tx.Where("operator_type=?", dto.OperatorType) // 操作类别（0其它 1后台用户 2手机端用户）
	}
	if len(dto.OperateName) > 0 {
		tx.Where("operate_name like %?%", dto.OperateName) // 操作人员
	}
	if len(dto.DeptName) > 0 {
		tx.Where("dept_name like %?%", dto.DeptName) // 部门名称
	}
	if len(dto.OperateUrl) > 0 {
		tx.Where("operate_url like %?%", dto.OperateUrl) // 请求URL
	}
	if len(dto.OperateIp) > 0 {
		tx.Where("operate_ip like %?%", dto.OperateIp) // 主机地址
	}
	if len(dto.OperateLocation) > 0 {
		tx.Where("operate_location like %?%", dto.OperateLocation) // 操作地点
	}

	if len(dto.Platform) > 0 {
		tx.Where("platform like %?%", dto.Platform) // 平台信息
	}
	if len(dto.Browser) > 0 {
		tx.Where("browser like %?%", dto.Browser) // 浏览器类型
	}
	if len(dto.Version) > 0 {
		tx.Where("version like %?%", dto.Version) // 浏览器版本
	}
	if len(dto.Os) > 0 {
		tx.Where("os like %?%", dto.Os) // 操作系统
	}

	if dto.Status != nil {
		tx.Where("status=?", dto.Status) // 操作状态(0:异常,正常)
	}

	tx.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&list)

	err := tx.Count(&total).Error
	return list, total, err
}
