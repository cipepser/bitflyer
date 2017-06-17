package main

import (
	"fmt"
	"strconv"
	"time"

	"./sdk"
)

const (
	// URL is a end point of bitflyer api.
	URL = "https://api.bitflyer.jp"
)

func main() {
	c, _ := sdk.NewClient(URL, "user", "passwd", nil)

	es := c.GetExecutions("FX_BTC_JPY", "", "", "")
	lid := es[0].ID

	for {
		for i := 0; i < len(es); i++ {
			fmt.Println(strconv.FormatFloat(es[len(es)-1-i].ID, 'f', 0, 64), ": ", es[len(es)-1-i].Price, "円 x ", es[len(es)-1-i].Size)
		}

		time.Sleep(5 * time.Second)
		lid = es[0].ID
		es = c.GetExecutions("FX_BTC_JPY", "", "", strconv.FormatFloat(lid, 'f', 0, 64))
	}

}
