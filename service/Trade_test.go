package service

import (
	"fmt"
	"testing"
	"github.com/BishengOpen/Bisheng-Golang-API-Demo/model"
)

func TestCreateOrder(t *testing.T) {

	var orderParams model.CreateOrderRequestParams
	orderParams.Side = "B"
	orderParams.Qty = "1.0"
	orderParams.Price = "1.0"
	orderParams.Symbol = "bsht_usdt"
	orderParams.Type = "L"

	fmt.Println("Create order: ", orderParams)
	createOrderReturn, err := CreateOrder(orderParams)
	if err != nil {
		t.Errorf("Create error: %v", err)
	}else{
		if createOrderReturn.Status == "ok" {
			fmt.Println("Create return: ", createOrderReturn.Data)
		} else {
			t.Errorf("Create error: %s", createOrderReturn.Error.ErrMsg)
		}
	}

}

func Test_getAccountBalance(t *testing.T)  {
	balance, err := GetAccountBalance("ETH_USDT")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(balance)
	}
}
