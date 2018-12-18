package main

import(
	"strconv"
	"fmt"

	"github.com/BishengOpen/Bisheng-Golang-API-Demo/service"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/model"

)

func main() {
	//test1()
	test2()
}

// 模拟下单流程
func test1(){
	//查询行情
	tiker, err := service.GetTicker("ETH_USDT")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n",tiker)
	}
	//查询持仓
	balance, err := service.GetAccountBalance("ETH_USDT")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n",balance)
	}

	//挂单
	var orderParams model.CreateOrderRequestParams
	orderParams.Side = "B"
	orderParams.Qty = "10"
	orderParams.Price = "0.1"
	orderParams.Symbol = "ETH_USDT"
	orderParams.Type = "L"

	fmt.Println("Create order: ", orderParams)
	createOrderReturn, err := service.CreateOrder(orderParams)
	if err != nil {
		fmt.Errorf("Create error: %v\n", err)
	}else{
		if createOrderReturn.Status == "ok" {
			fmt.Println("Create return: ", createOrderReturn.Data.OrderId)
		} else {
			fmt.Errorf("Create error: %s\n", createOrderReturn.Error.ErrMsg)
		}
	}

	//查询订单状态
	orderDetailReturn, err := service.GetOrderStatus(strconv.FormatInt(createOrderReturn.Data.OrderId, 10))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n",orderDetailReturn)
	}
}

// 模拟撤单流程
func test2(){
	//查询订单状态
	orderUndealReturn, err := service.GetOrderUnDeal("1","10")
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("%#v\n", orderUndealReturn)

		//撤单
		for _, order := range orderUndealReturn.Data{
			fmt.Println("orderId:", order.OrdNum)
			orderCancelReturn, err := service.CancelOrder(order.OrdNum)
			if err != nil {
				fmt.Println(err)
			} else {
				if orderCancelReturn.Status == "ok" {
					fmt.Println("Cancel return: ", orderCancelReturn.Data.OrderId)
				} else {
					fmt.Errorf("Cancel error: %s\n", orderCancelReturn.Error.ErrMsg)
				}
			}
		}
	}

}