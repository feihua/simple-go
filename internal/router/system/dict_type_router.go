package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DictTypeRouter 字典类型相关路由
func DictTypeRouter(r *gin.RouterGroup, b *a.DictTypeController) {

	r.POST("/dept/addDictType", b.CreateDictType)
	r.POST("/dept/deleteDictTypeByIds", b.DeleteDictTypeByIds)
	r.POST("/dept/updateDictType", b.UpdateDictType)
	r.POST("/dept/updateDictTypeStatus", b.UpdateDictTypeStatus)
	r.POST("/dept/queryDictTypeDetail", b.QueryDictTypeDetail)
	r.POST("/dept/queryDictTypeList", b.QueryDictTypeList)

}
