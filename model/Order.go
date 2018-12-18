package model

type CreateOrderRequestParams struct {
	Side      string `json:"side"` 		 // 订单类型，B/买单，S/卖单
	Qty       string `json:"qty"`        // 限价表示下单数量, 市价买单时表示买多少钱, 市价卖单时表示卖多少币
	Price     string `json:"price"`      // 下单价格, 市价单不传该参数
	Symbol    string `json:"mkt"`        // 交易对, btc_usdt, eth_btc......
	Type      string `json:"type"`       // 订单类型, M: 市价，L: 限价
}

type OrderIdData struct {
	OrderId int64   `json:"ordNum"`
}
type CreateOrderReturn struct {
	Status  string `json:"status"`
	Ts      int64   `json:"ts"`     // 发送时间
	Data    OrderIdData `json:"data"`
	Error   ErrorData  `json:"error"`
}

type DelOrderIdData struct {
	OrderId int64   `json:"delOrdNum"`
}
type CancelOrderReturn struct {
	Status  string `json:"status"`
	Ts      int64   `json:"ts"`     // 发送时间
	Data    DelOrderIdData `json:"data"`
	Error   ErrorData  `json:"error"`
}

type OrderDetailData struct {
	OrdNum     int64 `json:"ordNum"`      //订单编号
	PrdName    string `json:"prdName"`     // 基础产品名称
	OppoPrdName  string `json:"oppoPrdName"` // 计价产品
	OrdQty     string `json:"ordQty"`      // 申报数量
	OrdOppoQty string `json:"ordOppoQty"`  // 对价产品的数量，包含手续费
	OrdPrice   string `json:"ordPrice"`    // 申报价格
	OrdExeQty  string `json:"ordExeQty"`   // 成交数量
	OrdTime    int64  `json:"ordTime"`     // 申报时间，从1970-1-1以来的ns数
	Side       string `json:"side"`        // 买卖方向，只有B(买)和S(卖)
	OrdSts     string `json:"ordSts"`      // 订单状态，"I" //初始状态；"A" //全部成交；"E" //无效订单； "P" //部分成交，"C" //已经取消, "W" //待成交
}
type OrderDetailReturn struct {
	Status  string `json:"status"`
	Ts      int64   `json:"ts"`     // 发送时间
	Data    OrderDetailData `json:"data"`
	Error   ErrorData  `json:"error"`
}

type OrderUndealReturn struct {
	Status  string `json:"status"`
	Ts      int64   `json:"ts"`     // 发送时间
	Data    []OrderDetailData `json:"data"`
	Error   ErrorData  `json:"error"`
}