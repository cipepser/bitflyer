package usage

import "github.com/cipepser/bitflyer/src/sdk"

const (
	// URL is a end point of bitflyer api.
	URL = "https://api.bitflyer.jp"

// timeout = 10
)

func usage() {
	c, _ := sdk.NewClient(URL, "user", "passwd", nil)

	// ************** public API **************

	// 板情報の取得
	// b := c.GetBoard("FX_BTC_JPY")
	// fmt.Println(b)

	// Tickerの取得
	// t := c.GetTicker("FX_BTC_JPY")
	// fmt.Println(t)

	// 約定履歴の取得
	// es := c.GetExecutions("FX_BTC_JPY", "", "", "")
	// for _, e := range es {
	// 	fmt.Println(e)
	// 	fmt.Println("---------------")
	// }

	// ************** private API **************
	// 資産残高を取得
	// bs := c.GetBalances()
	// for _, b := range bs {
	// 	fmt.Println(b)
	// 	fmt.Println("---------------")
	// }

	// 証拠金の取得
	// col := c.GetCollateral()
	// fmt.Println(col)

	// 注文を出す
	// odr := sdk.ChildOrder{
	// 	ProductCode:    "FX_BTC_JPY",
	// 	ChildOrderType: "LIMIT",
	// 	Side:           "BUY",
	// 	Price:          200000,
	// 	Size:           0.001,
	// 	MinuteToExpire: 10,
	// 	TimeInForce:    "GTC",
	// }
	// res := c.SendNewOrder(odr)
	// fmt.Println("result: ", res)
	// res = c.SendNewOrder(odr)
	// fmt.Println("result: ", res)

	// 注文状態を取得
	// odrs := c.GetMyOrder("FX_BTC_JPY", "", "", "", "ACTIVE")
	// for _, odr := range odrs {
	// fmt.Println(odr.ChildOrderID)
	// }

	// 全注文をキャンセル
	// allcan := sdk.ChildOrderAllCanceled{
	// 	ProductCode: "FX_BTC_JPY",
	// }
	// res := c.CancelAllOrder(allcan)
	// fmt.Println("result: ", res)

	// 個別の注文をキャンセル
	// can := sdk.ChildOrderCanceled{
	// ProductCode:  "FX_BTC_JPY",
	// ChildOrderID: odrs[0].ChildOrderID,
	// }
	// res := c.CancelOrder(can)
	// fmt.Println("result: ", res)
}
