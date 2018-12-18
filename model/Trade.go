package model

type TradeData struct {
	Symbol   string `json:"mkt"`      //交易对名称
	TrdTime  int64  `json:"trdTime"`  //成交时间
	TrdPrice string `json:"trdPrice"` //成交价格
	Side     string `json:"side"`     //成交回报类型，对于当前用户来说是买单还是卖单，1/buy,2/sell
	TrdQty   string `json:"trdQty"`   //成交数量
	TrdNum   uint64 `json:"trdNum"`   //成交编号
	Fee      string `json:"fee"`      //交易费用
	TrdAmt   string `json:"trdAmt"`   //成交额
	Selfdeal int    `json:"selfDeal"` //1为自成交，0为他成交
	OrdNum   uint64 `json:"ordNum"`   //成交对应订单编号
}

type TradeReturn struct {
	Status  string      `json:"status"` // 请求状态, ok或者error
	Ts      int64       `json:"ts"`     // 发送时间
	Data    []TradeData `json:"data"`   // 成交记录
	Error   ErrorData  `json:"error"`
}
