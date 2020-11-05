package dict

import (
	"github.com/gin-gonic/gin"
	"simple-go/api/dict"
)

func DictUrl(r *gin.RouterGroup) {

	r.POST("/dict/add", dict.CreateDict)
	r.POST("/dict/list", dict.GetDictList)
	r.POST("/dict/update", dict.UpdateDict)
	r.POST("/dict/delete", dict.DeleteDictById)

}
