package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一返回
//参考：https://blog.csdn.net/upstream480/article/details/128168125
/*
Author: LiuFeiHua
Date: 2024/4/16 9:31
*/
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Error(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}

func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS.Code, SUCCESS.Msg, data)
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS.Code, msg, nil)
}

func Fail(c *gin.Context, err Response) {
	Result(c, err.Code, err.Msg, nil)
}

func FailWithMsg(c *gin.Context, err Response, msg string) {
	Result(c, err.Code, msg, nil)
}

var (
	SUCCESS         = Error(200, "success")
	NeedRedirect    = Error(301, "need redirect")
	InvalidArgs     = Error(400, "invalid params")
	Unauthorized    = Error(401, "unauthorized")
	Forbidden       = Error(403, "forbidden")
	NotFound        = Error(404, "not found")
	Conflict        = Error(409, "entry exist")
	TooManyRequests = Error(429, "too many requests")
	ResultError     = Error(500, "response error")
	DatabaseError   = Error(598, "database error")
	CSRFDetected    = Error(599, "csrf attack detected")

	ParamsError = Error(1000, "params error")

	UserError  = Error(2001, "username or password error")
	CodeExpire = Error(2002, "verification expire")
	CodeError  = Error(2003, "verification error")
	UserExist  = Error(2004, "user Exist")
	RoleError  = Error(3001, "role error")
	MenuError  = Error(4001, "menu error")
	DeptError  = Error(5001, "dept error")

	DictError = Error(6001, "dict error")

	LogError = Error(7001, "log error")
)
