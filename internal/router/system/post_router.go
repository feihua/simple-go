package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// PostRouter 岗位信息相关路由
func PostRouter(r *gin.RouterGroup, b *a.PostController) {

	r.POST("/dept/addPost", b.CreatePost)
	r.POST("/dept/deletePostByIds", b.DeletePostByIds)
	r.POST("/dept/updatePost", b.UpdatePost)
	r.POST("/dept/updatePostStatus", b.UpdatePostStatus)
	r.POST("/dept/queryPostDetail", b.QueryPostDetail)
	r.POST("/dept/queryPostList", b.QueryPostList)

}
