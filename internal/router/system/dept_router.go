package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// DeptRouter 部门相关路由
func DeptRouter(r *gin.RouterGroup, b *a.DeptController) {

	r.POST("/system/dept/addDept", b.CreateDept)
	r.POST("/system/dept/deleteDept", b.DeleteDeptByIds)
	r.POST("/system/dept/updateDept", b.UpdateDept)
	r.POST("/system/dept/updateDeptStatus", b.UpdateDeptStatus)
	r.POST("/system/dept/queryDeptDetail", b.QueryDeptDetail)
	r.POST("/system/dept/queryDeptList", b.QueryDeptList)

}
