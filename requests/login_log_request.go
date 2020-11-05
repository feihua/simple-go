package requests

type LoginLogRequest struct {
	ID       int64   `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
