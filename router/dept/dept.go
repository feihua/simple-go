package dept

import (
	"github.com/feihua/simple-go/controller/dept"
	"github.com/gin-gonic/gin"
)

func DeptUrl(r *gin.RouterGroup) {

	r.POST("/dept/add", dept.CreateDept)
	r.GET("/dept/list", dept.QueryDeptList)
	r.POST("/dept/update", dept.UpdateDept)
	r.POST("/dept/delete", dept.DeleteDeptById)

}
