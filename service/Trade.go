package service

import (
	"encoding/json"
	"strconv"

	"github.com/BishengOpen/Bisheng-Golang-API-Demo/model"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/utils"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/config"
)

// 此处只封装简单的3个API，其余可以自己添加

// 获取行情
// mkt: 交易对, btc_usdt, eth_btc......
// return: TickReturn对象
func GetTicker(strSymbol string) (model.TickerReturn, error) {
	tickerReturn := model.TickerReturn{}

	mapParams := make(map[string]string)
	mapParams["mkt"] = strSymbol

	strRequestUrl := config.TIKER
	//strUrl := config.BISHENG_URL + strRequestUrl

	jsonTickReturn, err := utils.ApiKeyGet(mapParams, strRequestUrl)
	if err != nil {
		return tickerReturn, err
	}
	err = json.Unmarshal([]byte(jsonTickReturn), &tickerReturn)
	if err != nil {
		return tickerReturn, err
	}

	return tickerReturn, nil
}

// 获取交易记录
// mkt: 交易对, btc_usdt, bsht_btc......
// pageNum: 开始页(从1开始）
// pageSize: 页大小（最大不超过100）
// startTime: 起始时间，unix时间，精确到秒，10位
// endTime: 起始时间，unix时间，精确到秒，10位
// return: TradeReturn对象
func GetTrade(strSymbol string, pageNum int, pageSize int, startTime int, endTime int) (model.TradeReturn, error) {
	tradeReturn := model.TradeReturn{}

	mapParams := make(map[string]string)
	mapParams["mkt"] = strSymbol
	mapParams["pageNum"] = strconv.Itoa(pageNum)
	mapParams["pageSize"] = strconv.Itoa(pageSize)
	mapParams["startTime"] = strconv.Itoa(startTime)
	mapParams["endTime"] = strconv.Itoa(endTime)

	strRequestUrl := config.RECORD_LIST

	jsonTradeReturn, err := utils.ApiKeyGet(mapParams, strRequestUrl)
	if err != nil {
		return tradeReturn, err
	}

	err = json.Unmarshal([]byte(jsonTradeReturn), &tradeReturn)

	if err != nil {
		return tradeReturn, err
	}
	return tradeReturn, nil
}


// 查询用户持仓
// prdNames: 产品名称，bsh：查询bsh持仓;bsh_eth_btc：查询bsh、eth、btc持仓,最多5个（产品名称大小写不敏感）
// return: BalanceReturn对象
func GetAccountBalance(strSymbol string) (model.BalanceReturn, error) {
	balanceReturn := model.BalanceReturn{}

	mapParams := make(map[string]string)
	mapParams["prdNames"] = strSymbol

	strRequest := config.BALANCE

	jsonBanlanceReturn, err := utils.ApiKeyGet(mapParams, strRequest)
	if err != nil {
		return balanceReturn, err
	}
	err = json.Unmarshal([]byte(jsonBanlanceReturn), &balanceReturn)
	if err != nil {
		return balanceReturn, err
	}
	return balanceReturn, nil
}

// 下单
// orderRequestParams: 下单信息
// return: CreateOrderReturn
func CreateOrder(orderRequestParams model.CreateOrderRequestParams) (model.CreateOrderReturn, error) {
	orderReturn := model.CreateOrderReturn{}

	mapParams := make(map[string]interface{})
	mapParams["side"] = orderRequestParams.Side
	mapParams["qty"] = orderRequestParams.Qty
	if 0 < len(orderRequestParams.Price) {
		mapParams["price"] = orderRequestParams.Price
	}else{
		mapParams["price"] = "0"
	}
	mapParams["mkt"] = orderRequestParams.Symbol
	mapParams["type"] = orderRequestParams.Type

	strRequest := config.ORDER_CREATE
	jsonOrderReturn, err := utils.ApiKeyPost(mapParams, strRequest)
	if err != nil {
		return orderReturn, err
	}

	err = json.Unmarshal([]byte(jsonOrderReturn), &orderReturn)
	if err != nil {
		return orderReturn, err
	}
	return orderReturn, nil
}

// 撤单
// strOrderID: 订单ID
// return: CancelOrderReturn对象
func CancelOrder(orderID int64) (model.CancelOrderReturn, error) {
	cancelOrderlReturn := model.CancelOrderReturn{}

	mapParams := make(map[string]interface{})
	mapParams["ordNum"] = orderID

	strRequest := config.ORDER_CANCEL
	jsonOrderReturn, err := utils.ApiKeyPost(mapParams, strRequest)
	if err != nil {
		return cancelOrderlReturn, err
	}

	err = json.Unmarshal([]byte(jsonOrderReturn), &cancelOrderlReturn)
	if err != nil {
		return cancelOrderlReturn, err
	}
	return cancelOrderlReturn, nil
}

// 查询订单状态
// strOrderID: 订单ID
// return: TradeReturn对象
func GetOrderStatus(strOrderID string) (model.OrderDetailReturn, error) {
	orderStatusReturn := model.OrderDetailReturn{}

	mapParams := make(map[string]string)
	mapParams["ordNum"] = strOrderID

	strRequestUrl := config.ORDER_STATUS

	jsonOrderStatusReturn, err := utils.ApiKeyGet(mapParams, strRequestUrl)
	if err != nil {
		return orderStatusReturn, err
	}

	err = json.Unmarshal([]byte(jsonOrderStatusReturn), &orderStatusReturn)

	if err != nil {
		return orderStatusReturn, err
	}
	return orderStatusReturn, nil
}

// 查询所有未成交订单
// pageNum: 页(从1开始)
// pageSize: 页大小（最大不超过100）
// return: OrderUndealReturn对象
func GetOrderUnDeal(pageNum string, pageSize string) (model.OrderUndealReturn, error) {
	orderUndealReturn := model.OrderUndealReturn{}

	mapParams := make(map[string]string)
	mapParams["pageNum"] = pageNum
	mapParams["pageSize"] = pageSize

	strRequestUrl := config.ORDER_UNDEAL

	jsonOrderUndealReturn, err := utils.ApiKeyGet(mapParams, strRequestUrl)
	if err != nil {
		return orderUndealReturn, err
	}

	err = json.Unmarshal([]byte(jsonOrderUndealReturn), &orderUndealReturn)

	if err != nil {
		return orderUndealReturn, err
	}
	return orderUndealReturn, nil
}
