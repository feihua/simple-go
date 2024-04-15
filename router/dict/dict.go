package dict

import (
	"github.com/feihua/simple-go/controller/dict"
	"github.com/gin-gonic/gin"
)

func DictUrl(r *gin.RouterGroup) {

	r.POST("/dict/add", dict.CreateDict)
	r.GET("/dict/list", dict.QueryDictList)
	r.POST("/dict/update", dict.UpdateDict)
	r.POST("/dict/delete", dict.DeleteDictById)

}
