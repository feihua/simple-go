package user

import (
	"github.com/feihua/simple-go/dto"
	"github.com/feihua/simple-go/services/user"
	"github.com/feihua/simple-go/vo/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {

	//校验通过，返回请求参数
	c.JSON(http.StatusOK, gin.H{"status": "ok", "type": "account", "currentAuthority": "admin"})
}

func GetUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"name": "liufeihua", "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"})
}

func GetUser(c *gin.Context) {

	//实例化一个TestRequest结构体，用于接收参数
	userRequest := requests.UserRequest{}

	//接收请求参数
	err := c.ShouldBind(&userRequest)

	//判断参数校验是否通过，如果不通过，把错误返回给前端
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service user.UserContract = &user.UserService{}

	result, _ := service.QueryUserList(1, 10)

	//校验通过，返回请求参数
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func CreateUser(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service user.UserContract = &user.UserService{}

	u := dto.UserDto{
		Mobile:   userRequest.Mobile,
		UserName: userRequest.UserName,
		Password: userRequest.Password,
		StatusId: userRequest.StatusId,
		Sort:     userRequest.Sort,
		Remark:   userRequest.Remark,
	}
	result := service.CreateUser(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func QueryUserList(c *gin.Context) {
	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	current := c.DefaultQuery("current", "1")
	pageNum, _ := strconv.Atoi(current)

	pageSize := c.DefaultQuery("pageSize", "10")
	size, _ := strconv.Atoi(pageSize)

	var service user.UserContract = &user.UserService{}

	result, total := service.QueryUserList(pageNum, size)

	println(total)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true, "current": current, "total": total, "pageSize": pageSize})
}

func UpdateUser(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service user.UserContract = &user.UserService{}

	u := dto.UserDto{
		Id:       userRequest.Id,
		Mobile:   userRequest.Mobile,
		UserName: userRequest.UserName,
		Password: userRequest.Password,
		StatusId: userRequest.StatusId,
		Sort:     userRequest.Sort,
		Remark:   userRequest.Remark,
	}
	result := service.UpdateUser(u)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteUserById(c *gin.Context) {

	userRequest := requests.UserRequest{}
	err := c.ShouldBind(&userRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service user.UserContract = &user.UserService{}

	result := service.DeleteUserById(userRequest.Id)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdateUserRole(c *gin.Context) {

	req := requests.UserRoleRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var service user.UserContract = &user.UserService{}

	err = service.UpdateUserRole(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "成功"})
}
