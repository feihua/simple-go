package dept

import (
	"github.com/feihua/simple-go/controller/dept"
	"github.com/gin-gonic/gin"
)

// DeptRouter 部门相关路由
func DeptRouter(r *gin.RouterGroup) {

	controller := dept.NewDeptController()
	r.POST("/dept/addDept", controller.CreateDept)
	r.GET("/dept/queryDeptList", controller.QueryDeptList)
	r.POST("/dept/updateDept", controller.UpdateDept)
	r.GET("/dept/deleteDeptByIds", controller.DeleteDeptByIds)

}
