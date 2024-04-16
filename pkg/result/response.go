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
	UserError   = Error(5001, "username or password error")
	CodeExpire  = Error(5002, "verification expire")
	CodeError   = Error(5003, "verification error")
	UserExist   = Error(5004, "user Exist")

	DeptError = Error(6001, "dept error")

	DictError = Error(7001, "dict error")
)
