package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"time"
)

func main() {
	u := "https://api.bitflyer.jp/"
	c, _ := NewClient(u, "user", "passwd", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body io.Reader
	req, err := c.newRequest(ctx, "GET", "v1/board", body)
	if err != nil {
		log.Fatal(err)
	}

	values := url.Values{}
	values.Add("product_code", "FX_BTC_JPY")
	req.URL.RawQuery = values.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	board := Board{}
	err = decodeBody(resp, &board)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(board)

}
