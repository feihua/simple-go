package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DictTypeRouter 字典类型相关路由
func DictTypeRouter(r *gin.RouterGroup, b *a.DictTypeController) {

	r.POST("/system/dictType/addDictType", b.CreateDictType)
	r.POST("/system/dictType/deleteDictTypeByIds", b.DeleteDictTypeByIds)
	r.POST("/system/dictType/updateDictType", b.UpdateDictType)
	r.POST("/system/dictType/updateDictTypeStatus", b.UpdateDictTypeStatus)
	r.POST("/system/dictType/queryDictTypeDetail", b.QueryDictTypeDetail)
	r.POST("/system/dictType/queryDictTypeList", b.QueryDictTypeList)

}
