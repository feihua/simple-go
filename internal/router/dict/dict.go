package dict

import (
	"github.com/feihua/simple-go/internal/controller/dict"
	"github.com/gin-gonic/gin"
)

func DictRouter(r *gin.RouterGroup, dictController *dict.DictController) {

	r.POST("/dict/addDict", dictController.CreateDict)
	r.POST("/dict/queryDictList", dictController.QueryDictList)
	r.POST("/dict/updateDict", dictController.UpdateDict)
	r.POST("/dict/deleteDictByIds", dictController.DeleteDictByIds)

}
