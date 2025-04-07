package responses

// LoginResp 登录响应
type LoginResp struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}
