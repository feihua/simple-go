package dept

import (
	"github.com/feihua/simple-go/controller"
	"github.com/gin-gonic/gin"
)

// DeptRouter 部门相关路由
func DeptRouter(r *gin.RouterGroup) {

	deptController := controller.C.DeptController
	r.POST("/dept/addDept", deptController.CreateDept)
	r.GET("/dept/queryDeptList", deptController.QueryDeptList)
	r.POST("/dept/updateDept", deptController.UpdateDept)
	r.GET("/dept/deleteDeptByIds", deptController.DeleteDeptByIds)

}
