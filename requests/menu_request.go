package requests

type MenuRequest struct {
	ID       int   `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
