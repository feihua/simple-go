package dto

type LogDto struct {
	UserName       string  `json:"userName"`       //用户名
	Operation      string  `json:"operation"`      //用户操作
	Method         string  `json:"method"`         //请求方法
	RequestParams  string  `json:"requestParams"`  //请求参数
	ResponseParams *string `json:"responseParams"` //响应参数
	OperationTime  int64   `json:"operationTime"`  //执行时长(毫秒)
	Ip             *string `json:"ip"`             //IP地址
}
