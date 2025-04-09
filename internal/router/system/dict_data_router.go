package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DictDataRouter 字典数据相关路由
func DictDataRouter(r *gin.RouterGroup, b *a.DictDataController) {

	r.POST("/dept/addDictData", b.CreateDictData)
	r.POST("/dept/deleteDictDataByIds", b.DeleteDictDataByIds)
	r.POST("/dept/updateDictData", b.UpdateDictData)
	r.POST("/dept/updateDictDataStatus", b.UpdateDictDataStatus)
	r.POST("/dept/queryDictDataDetail", b.QueryDictDataDetail)
	r.POST("/dept/queryDictDataList", b.QueryDictDataList)

}
