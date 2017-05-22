package main

import (
	"fmt"

	"github.com/cipepser/bitflyer/src/sdk"
)

func main() {
	fmt.Println(sdk.GetCollateral())

	fmt.Println(sdk.GetBoard("FX_BTC_JPY"))
}
