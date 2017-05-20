package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"./sdk"
)

// "github.com/cipepser/bitflyer/src/sdk"

func main() {
	u := "https://api.bitflyer.jp"
	c, _ := sdk.NewClient(u, "user", "passwd", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body io.Reader

	// 	body = "{
	//     product_code: 'BTC_JPY',
	//     child_order_type: 'LIMIT',
	//     side: 'BUY',
	//     price: 30000,
	//     size: 0.1
	// }"

	req, err := c.NewRequest(ctx, "GET", "/v1/me/getcollateral", body)
	if err != nil {
		log.Fatal(err)
	}

	// values := url.Values{}
	// values.Add("product_code", "FX_BTC_JPY")
	// req.URL.RawQuery = values.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// board := sdk.Board{}
	// err = sdk.DecodeBody(resp, &board)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(board)

	collateral := sdk.Collateral{}
	// board := sdk.Board{}
	err = sdk.DecodeBody(resp, &collateral)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collateral)

	// var body io.Reader

	// fmt.Println(body)

}
