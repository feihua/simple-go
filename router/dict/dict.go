package dict

import (
	"github.com/feihua/simple-go/controller/dict"
	"github.com/gin-gonic/gin"
)

func DictUrl(r *gin.RouterGroup) {

	controller := dict.NewDictController()
	r.POST("/dict/addDict", controller.CreateDict)
	r.GET("/dict/queryDictList", controller.QueryDictList)
	r.POST("/dict/updateDict", controller.UpdateDict)
	r.GET("/dict/deleteDictByIds", controller.DeleteDictByIds)

}
