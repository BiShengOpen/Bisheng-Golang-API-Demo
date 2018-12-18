package model

type Balance struct {
	PrdName string `json:"prdName"`    //币种
	Qty     string `json:"qty"`		   //持仓数量
	QtyRa   string `json:"qtyRa"`      //冻结数量
}

type BalanceReturn struct {
	Status  string  `json:"status"` // 请求状态
	Ts      int64   `json:"ts"`     // 发送时间
	Data    []Balance `json:"data"` // 账户余额
	Error   ErrorData  `json:"error"`
}
