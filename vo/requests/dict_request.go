package requests

type DictRequest struct {
	Id          int64   `json:"id"`          //编号
	Value       string  `json:"value"`       //数据值
	Label       string  `json:"label"`       //标签名
	Type        string  `json:"type"`        //类型
	Description string  `json:"description"` //描述
	Sort        float64 `json:"sort"`        //排序（升序）
	Remark      *string `json:"remark"`      //备注
}
type DeleteDictRequest struct {
	Ids []int64 `json:"ids"` //编号
}
