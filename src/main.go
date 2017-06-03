package main

import (
	"./myUtil"
	"./sdk"
)

const (
	// URL is a end point of bitflyer api.
	URL = "https://api.bitflyer.jp"

// timeout = 10
)

func main() {
	c, _ := sdk.NewClient(URL, "user", "passwd", nil)
	es := c.GetExecutions("FX_BTC_JPY", "", "", "")

	x := make([]float64, len(es))
	for i, e := range es {
		x[i] = e.Price
	}

	// x := make([]float64, 10)
	// y := make([]float64, 10)
	// for i := 0; i < len(x); i++ {
	// x[i] = (float64(i) - 3) * (float64(i) - 7) * float64(i)
	// x[i] = float64(i)
	// y[i] = (float64(i) - 3) * (float64(i) - 7) * float64(i)
	// }

	// fmt.Println(x)

	// myUtil.MyScatter(x, y)
	myUtil.MySingleScatter(x)

}
