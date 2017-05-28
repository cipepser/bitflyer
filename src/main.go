package main

import (
	"io"

	"./sdk"
)

func test(s string, b io.Reader) {
	return
}

const (
	// 	// URL is a end point of bitflyer api.
	URL = "https://api.bitflyer.jp"

// timeout = 10
)

func main() {
	c, _ := sdk.NewClient(URL, "user", "passwd", nil)
	// b := c.GetBoard("FX_BTC_JPY")
	// fmt.Println(b)

	// col := c.GetCollateral()
	// fmt.Println(col)

	// TODO: wi-fiあるところでテスト
	// odr := sdk.ChildOrder{
	// 	ProductCode:    "FX_BTC_JPY",
	// 	ChildOrderType: "LIMIT",
	// 	Side:           "BUY",
	// 	Price:          200000,
	// 	Size:           0.001,
	// 	MinuteToExpire: 1,
	// 	TimeInForce:    "GTC",
	// }
	//
	// res := c.SendNewOrder(odr)
	// fmt.Println("result: ", res)
	//
	// odrs := c.GetMyOrder("FX_BTC_JPY", "", "", "", "ACTIVE")
	// fmt.Println(len(odrs))

}
