package requests

type SysLogRequest struct {
	Id            int64  `json:"id"`            //主键
	UserName      string `json:"userName"`      //用户名
	Operation     string `json:"operation"`     //用户操作
	Method        string `json:"method"`        //请求方法
	Params        string `json:"params"`        //请求参数
	OperationTime int64  `json:"operationTime"` //执行时长(毫秒)
	Ip            string `json:"ip"`            //IP地址
}
type DeleteLogRequest struct {
	Ids []int64 `json:"ids"` //编号
}
