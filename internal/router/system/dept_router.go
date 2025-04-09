package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DeptRouter 部门相关路由
func DeptRouter(r *gin.RouterGroup, b *a.DeptController) {

	r.POST("/dept/addDept", b.CreateDept)
	r.POST("/dept/deleteDeptByIds", b.DeleteDeptByIds)
	r.POST("/dept/updateDept", b.UpdateDept)
	r.POST("/dept/updateDeptStatus", b.UpdateDeptStatus)
	r.POST("/dept/queryDeptDetail", b.QueryDeptDetail)
	r.POST("/dept/queryDeptList", b.QueryDeptList)

}
