package system

import (
	a "github.com/feihua/simple-go/internal/controller/system"
	"github.com/gin-gonic/gin"
)

// PostRouter 岗位信息相关路由
func PostRouter(r *gin.RouterGroup, b *a.PostController) {

	r.POST("/system/post/addPost", b.CreatePost)
	r.POST("/system/post/deletePost", b.DeletePostByIds)
	r.POST("/system/post/updatePost", b.UpdatePost)
	r.POST("/system/post/updatePostStatus", b.UpdatePostStatus)
	r.POST("/system/post/queryPostDetail", b.QueryPostDetail)
	r.POST("/system/post/queryPostList", b.QueryPostList)

}
