package operate_log

import (
	"errors"
	"github.com/feihua/simple-go/internal/dao/system"
	d "github.com/feihua/simple-go/internal/dto/system"
	"github.com/feihua/simple-go/pkg/utils"
)

// OperateLogServiceImpl 操作日志记录操作实现
type OperateLogServiceImpl struct {
	Dao *system.OperateLogDao
}

func NewOperateLogServiceImpl(dao *system.OperateLogDao) OperateLogService {
	return &OperateLogServiceImpl{
		Dao: dao,
	}
}

// CreateOperateLog 添加操作日志记录
func (s *OperateLogServiceImpl) CreateOperateLog(dto d.AddOperateLogDto) error {
	return s.Dao.CreateOperateLog(dto)
}

// DeleteOperateLogByIds 删除操作日志记录
func (s *OperateLogServiceImpl) DeleteOperateLogByIds(ids []int64) error {
	return s.Dao.DeleteOperateLogByIds(ids)
}

// QueryOperateLogDetail 查询操作日志记录详情
func (s *OperateLogServiceImpl) QueryOperateLogDetail(dto d.QueryOperateLogDetailDto) (*d.QueryOperateLogListDtoResp, error) {
	item, err := s.Dao.QueryOperateLogDetail(dto)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("操作日志记录不存在")
	}

	return &d.QueryOperateLogListDtoResp{
		Id:              item.Id,                           // 日志主键
		Title:           item.Title,                        // 模块标题
		BusinessType:    item.BusinessType,                 // 业务类型（0其它 1新增 2修改 3删除）
		Method:          item.Method,                       // 方法名称
		RequestMethod:   item.RequestMethod,                // 请求方式
		OperatorType:    item.OperatorType,                 // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     item.OperateName,                  // 操作人员
		DeptName:        item.DeptName,                     // 部门名称
		OperateUrl:      item.OperateUrl,                   // 请求URL
		OperateIp:       item.OperateIp,                    // 主机地址
		OperateLocation: item.OperateLocation,              // 操作地点
		OperateParam:    item.OperateParam,                 // 请求参数
		JsonResult:      item.JsonResult,                   // 返回参数
		Platform:        item.Platform,                     // 平台信息
		Browser:         item.Browser,                      // 浏览器类型
		Version:         item.Version,                      // 浏览器版本
		Os:              item.Os,                           // 操作系统
		Arch:            item.Arch,                         // 体系结构信息
		Engine:          item.Engine,                       // 渲染引擎信息
		EngineDetails:   item.EngineDetails,                // 渲染引擎详细信息
		Extra:           item.Extra,                        // 其他信息（可选）
		Status:          item.Status,                       // 操作状态(0:异常,正常)
		ErrorMsg:        item.ErrorMsg,                     // 错误消息
		OperateTime:     utils.TimeToStr(item.OperateTime), // 操作时间
		CostTime:        item.CostTime,                     // 消耗时间
	}, nil

}

// QueryOperateLogList 查询操作日志记录列表
func (s *OperateLogServiceImpl) QueryOperateLogList(dto d.QueryOperateLogListDto) ([]*d.QueryOperateLogListDtoResp, int64, error) {
	result, i, err := s.Dao.QueryOperateLogList(dto)

	if err != nil {
		return nil, 0, err
	}

	list := make([]*d.QueryOperateLogListDtoResp, 0)

	for _, item := range result {
		resp := &d.QueryOperateLogListDtoResp{
			Id:              item.Id,                           // 日志主键
			Title:           item.Title,                        // 模块标题
			BusinessType:    item.BusinessType,                 // 业务类型（0其它 1新增 2修改 3删除）
			Method:          item.Method,                       // 方法名称
			RequestMethod:   item.RequestMethod,                // 请求方式
			OperatorType:    item.OperatorType,                 // 操作类别（0其它 1后台用户 2手机端用户）
			OperateName:     item.OperateName,                  // 操作人员
			DeptName:        item.DeptName,                     // 部门名称
			OperateUrl:      item.OperateUrl,                   // 请求URL
			OperateIp:       item.OperateIp,                    // 主机地址
			OperateLocation: item.OperateLocation,              // 操作地点
			OperateParam:    item.OperateParam,                 // 请求参数
			JsonResult:      item.JsonResult,                   // 返回参数
			Platform:        item.Platform,                     // 平台信息
			Browser:         item.Browser,                      // 浏览器类型
			Version:         item.Version,                      // 浏览器版本
			Os:              item.Os,                           // 操作系统
			Arch:            item.Arch,                         // 体系结构信息
			Engine:          item.Engine,                       // 渲染引擎信息
			EngineDetails:   item.EngineDetails,                // 渲染引擎详细信息
			Extra:           item.Extra,                        // 其他信息（可选）
			Status:          item.Status,                       // 操作状态(0:异常,正常)
			ErrorMsg:        item.ErrorMsg,                     // 错误消息
			OperateTime:     utils.TimeToStr(item.OperateTime), // 操作时间
			CostTime:        item.CostTime,                     // 消耗时间
		}

		list = append(list, resp)
	}

	return list, i, nil
}
