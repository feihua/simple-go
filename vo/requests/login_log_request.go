package requests

type LoginLogRequest struct {
	Id       int64  `json:"id"`       //主键
	UserName string `json:"userName"` //用户名
	Ip       string `json:"ip"`       //IP地址
}
type DeleteLoginLogRequest struct {
	Ids []int64 `json:"ids"` //编号
}
