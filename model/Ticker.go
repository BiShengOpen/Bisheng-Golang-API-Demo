package model

type Record struct {
	Price float64 `json:"ordPrice"`
	Qty   float64 `json:"ordQty"`
}

type Ticker struct {
	Id     int      `json:"id"`           //K线id
	Ts     int64    `json:"ts"`           //时间戳
	Close  float64  `json:"close"`        //收盘价，现在为开盘价
	Open   float64  `json:"open"`         //开盘价
	Low    float64  `json:"low"`          //最低价
	High   float64  `json:"high"`         //最高价
	Amount float64  `json:"todayDealQty"` //今日成交量
	Amt24  float64  `json:"amt24"`        //24小时成交额
	Vol    float64  `json:"vol"`          //24小时成交量
	Ask    []Record `json:"sell"`         //卖一eg:[卖1价,卖1量]
	Bid    []Record `json:"buy"`          //买一  eg:[买1价,买1量]
}

type TickerReturn struct {
	Status  string `json:"status"` // 请求处理结果
	Ts      int64  `json:"ts"`     // 响应生成时间点
	Data    Ticker `json:"data"`   // 聚合数据
	Error   ErrorData  `json:"error"`
}