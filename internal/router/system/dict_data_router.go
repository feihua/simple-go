package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DictDataRouter 字典数据相关路由
func DictDataRouter(r *gin.RouterGroup, b *a.DictDataController) {

	r.POST("/system/dictData/addDictData", b.CreateDictData)
	r.POST("/system/dictData/deleteDictData", b.DeleteDictDataByIds)
	r.POST("/system/dictData/updateDictData", b.UpdateDictData)
	r.POST("/system/dictData/updateDictDataStatus", b.UpdateDictDataStatus)
	r.POST("/system/dictData/queryDictDataDetail", b.QueryDictDataDetail)
	r.POST("/system/dictData/queryDictDataList", b.QueryDictDataList)

}
