package dict

import (
	"github.com/feihua/simple-go/controller"
	"github.com/gin-gonic/gin"
)

func DictRouter(r *gin.RouterGroup) {

	dictController := controller.C.DictController
	r.POST("/dict/addDict", dictController.CreateDict)
	r.GET("/dict/queryDictList", dictController.QueryDictList)
	r.POST("/dict/updateDict", dictController.UpdateDict)
	r.GET("/dict/deleteDictByIds", dictController.DeleteDictByIds)

}
